package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

func init() {
	eh.RegisterAggregate(func(id uuid.UUID) eh.Aggregate {
		return &Aggregate{
			AggregateBase: events.NewAggregateBase(AggregateType, id),
		}
	})
}

// AggregateType specifies the appregate type
const AggregateType = eh.AggregateType("payment")

// Aggregate represents the payment aggregate
type Aggregate struct {
	*events.AggregateBase

	created    bool
	deleted    bool
	attributes PaymentAttributes
}

// TimeNow is version of Time.Now for mocking
var TimeNow = time.Now

// HandleCommand is the command handler
func (a *Aggregate) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd.(type) {
	case *Create:
		// Payment can only be created once
		if a.created {
			return errors.New("payment already created")
		}
	case *Delete:
		// To delete it must be created
		if !a.created {
			return errors.New("payment must be created to delete")
		}
		// Can only be deleted once
		if a.deleted {
			return errors.New("payment already deleted")
		}
	}

	switch cmd := cmd.(type) {
	case *Create:
		if cmd.Attributes.Amount <= 0 {
			return fmt.Errorf("could not created a payment with an amount of 0 or less")
		}
		//TODO: Add propery validation of the command, i.e. check account numbers etc

		a.StoreEvent(Created, &CreatedData{
			Attributes: cmd.Attributes,
		}, TimeNow())
	case *Update:
		if cmd.Attributes.Amount <= 0 {
			return fmt.Errorf("could not update a payment with an amount of 0 or less")
		}
		//TODO: Add propery validation of the command, i.e. check account numbers etc

		a.StoreEvent(Updated, &UpdatedData{
			Attributes: cmd.Attributes,
		}, TimeNow())
	case *Delete:
		a.StoreEvent(Deleted, &DeletedData{
			Reason: cmd.Reason,
		}, TimeNow())
	default:
		return fmt.Errorf("could not handle command: %s", cmd.CommandType())
	}
	return nil
}

// ApplyEvent applies events to the aggregate
func (a *Aggregate) ApplyEvent(ctx context.Context, event eh.Event) error {
	switch event.EventType() {
	case Created:
		a.created = true
		data, ok := event.Data().(*CreatedData)
		if !ok {
			return errors.New("invalid event data")
		}
		a.attributes = data.Attributes
	case Updated:
		data, ok := event.Data().(*UpdatedData)
		if !ok {
			return errors.New("invalid event data")
		}
		a.attributes = data.Attributes
	case Deleted:
		a.deleted = true
	default:
		return fmt.Errorf("could not apply event: %s", event.EventType())
	}
	return nil
}
