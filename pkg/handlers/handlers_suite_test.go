package handlers_test

import (
	"net/http"
	"testing"
	"time"
	"context"

	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/memory"
	repo "github.com/looplab/eventhorizon/repo/memory"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/waiter"
	. "github.com/onsi/ginkgo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/richardcase/paymentssvc/pkg/config"
	"github.com/richardcase/paymentssvc/pkg/domain"
	"github.com/richardcase/paymentssvc/pkg/gen/restapi"
	"github.com/richardcase/paymentssvc/pkg/handlers"
	"github.com/richardcase/paymentssvc/pkg/testutils"
)

func TestSuite(t *testing.T) {
	testutils.RegisterAndRun(t)
}

var _ = BeforeSuite(func() {
	domain.TimeNow = func() time.Time {
		return time.Date(2019, time.January, 10, 23, 0, 0, 0, time.UTC)
	}
})

func createDomainAttributes(amount float32) domain.PaymentAttributes {
	return domain.PaymentAttributes{
		Amount:        amount,
		Currency:      "GBP",
		PaymentScheme: "FPS",
		PaymentType:   "Credit",
		BeneficiaryParty: domain.Party{
			AccountName:   "Test1",
			AccountNumber: "12345",
		},
		DebtorParty: domain.Party{
			AccountName:   "Test2",
			AccountNumber: "56789",
		},
	}
}

func createTestConfig() config.Config {
	logger := logrus.WithFields(logrus.Fields{"test": "true"})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter((&logrus.TextFormatter{}))

	eventStore := eventstore.NewEventStore()
	eventBus := eventbus.NewEventBus(nil)
	repo := repo.NewRepo()

	return config.Config{
		EventStore: eventStore,
		EventBus:   eventBus,
		Repo:       repo,
		Logger:     logger,
	}
}

func createTestHandler(config config.Config) (http.Handler, *handlers.Handlers, error) {
	api, err := restapi.GetAPI()
	if err != nil {
		return nil, nil, errors.Wrap(err, "creating REST api")
	}

	h, err := handlers.New(config, api)
	if err != nil {
		return nil, nil, errors.Wrap(err, "creating new handler")
	}

	restHandler, err := restapi.GetAPIHandler(api)
	if err != nil {
		return nil, nil, errors.Wrap(err, "creating REST handler")
	}

	return restHandler, h, nil
}

func waitForEvent(bus eh.EventBus, eventType eh.EventType) {
	waiter := waiter.NewEventHandler()
	bus.AddObserver(eh.MatchEvent(eventType), waiter)
	l := waiter.Listen(nil)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	l.Wait(ctx)
}
