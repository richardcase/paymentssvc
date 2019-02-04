package domain

import (
	"context"
	"errors"
	"fmt"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventhandler/projector"
)

// Projector will project the events of a payment into a Payment read model.
type Projector struct{}

// ProjectorType return the type of prjector
func (p *Projector) ProjectorType() projector.Type {
	return projector.Type(string(AggregateType) + "_projector")
}

// Project performs the projection of events into the read model
func (p *Projector) Project(ctx context.Context,
	event eh.Event, entity eh.Entity) (eh.Entity, error) {

	model, ok := entity.(*Payment)
	if !ok {
		return nil, errors.New("model is of incorrect type")
	}

	switch event.EventType() {
	case Created:
		model.ID = event.AggregateID()
		model.CreatedAt = TimeNow()

		data, ok := event.Data().(*CreatedData)
		if !ok {
			return nil, errors.New("invalid event data")
		}
		model.Attributes = data.Attributes
	case Updated:
		data, ok := event.Data().(*UpdatedData)
		if !ok {
			return nil, errors.New("invalid event data")
		}
		model.Attributes = data.Attributes
	case Deleted:
		return nil, nil
	default:
		return model, fmt.Errorf("could not project event: %s", event.EventType())
	}

	model.Version++
	model.UpdatedAt = TimeNow()
	return model, nil
}
