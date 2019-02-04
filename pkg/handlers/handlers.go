package handlers

import (
	"context"
	"fmt"
	"github.com/richardcase/paymentssvc/pkg/gen/models"
	"sync"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/commandhandler/aggregate"
	"github.com/looplab/eventhorizon/eventhandler/projector"
	"github.com/looplab/eventhorizon/repo/version"

	"github.com/richardcase/paymentssvc/pkg/config"
	"github.com/richardcase/paymentssvc/pkg/domain"
	"github.com/richardcase/paymentssvc/pkg/gen/restapi/operations"
	"github.com/richardcase/paymentssvc/pkg/gen/restapi/operations/payment"
	"github.com/richardcase/paymentssvc/pkg/gen/restapi/operations/probes"
)

// Handlers contains the handler implementation for the API
type Handlers struct {
	config.Config

	EventBus       eh.EventBus
	CommandHandler eh.CommandHandler
	Repo           eh.ReadWriteRepo

	readyMu sync.RWMutex
	ready   bool
}

// New creates a new handler
func New(conf config.Config, api *operations.PaymentsAPI) (*Handlers, error) {
	handlers := &Handlers{
		Config: conf,
	}

	api.PaymentAddPaymentHandler = payment.AddPaymentHandlerFunc(handlers.AddPayment)
	api.PaymentDeletePaymentHandler = payment.DeletePaymentHandlerFunc(handlers.DeletePayment)
	api.PaymentGetPaymentByIDHandler = payment.GetPaymentByIDHandlerFunc(handlers.GetPaymentByID)
	api.PaymentListPaymentsHandler = payment.ListPaymentsHandlerFunc(handlers.ListPayments)
	api.PaymentUpdatePaymentHandler = payment.UpdatePaymentHandlerFunc(handlers.UpdatePayment)
	api.ProbesGetLiveHandler = probes.GetLiveHandlerFunc(handlers.GetLive)
	api.ProbesGetReadyHandler = probes.GetReadyHandlerFunc(handlers.GetReady)

	api.ServerShutdown = func() {}

	// Create the aggregate repository.
	aggregateStore, err := events.NewAggregateStore(conf.EventStore, conf.EventBus)
	if err != nil {
		return nil, fmt.Errorf("could not create aggregate store: %s", err)
	}

	// Create the aggregate command handler.
	aggregateCommandHandler, err := aggregate.NewCommandHandler(domain.AggregateType, aggregateStore)
	if err != nil {
		return nil, fmt.Errorf("could not create command handler: %s", err)
	}

	// Create a tiny logging middleware for the command handler.
	commandHandlerLogger := func(h eh.CommandHandler) eh.CommandHandler {
		return eh.CommandHandlerFunc(func(ctx context.Context, cmd eh.Command) error {
			conf.Logger.Debugf("CMD %#v", cmd)
			return h.HandleCommand(ctx, cmd)
		})
	}
	commandHandler := eh.UseCommandHandlerMiddleware(aggregateCommandHandler, commandHandlerLogger)

	paymentRepo := version.NewRepo(conf.Repo)
	// Create the read model projector.
	projector := projector.NewEventHandler(&domain.Projector{}, paymentRepo)
	projector.SetEntityFactory(func() eh.Entity { return &domain.Payment{} })
	conf.EventBus.AddHandler(eh.MatchAnyEventOf(
		domain.Created,
		domain.Deleted,
		domain.Updated,
	), projector)

	handlers.EventBus = conf.EventBus
	handlers.CommandHandler = commandHandler
	handlers.Repo = paymentRepo

	return handlers, nil
}

// SetReady will set the ready prode to true
func (h *Handlers) SetReady() {
	h.readyMu.Lock()
	defer h.readyMu.Unlock()

	h.ready = true
}

// SetNotReady starts the readiness probe indicating we are not ready
func (h *Handlers) SetNotReady() {
	h.readyMu.Lock()
	defer h.readyMu.Unlock()
	h.ready = false
}

func (h *Handlers) createError(message string) *models.Error {
	return &models.Error{
		Message: &message,
	}
}
