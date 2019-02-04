package domain

import (
	eh "github.com/looplab/eventhorizon"
)

const (
	// Created is the event type after a payment is created
	Created = eh.EventType("payment:created")
	// Deleted is the event type after a payment is deleted
	Deleted = eh.EventType("payment:deleted")
	// Updated is the event type after a payment is updated
	Updated = eh.EventType("payment:updated")
)

func init() {
	eh.RegisterEventData(Created, func() eh.EventData {
		return &CreatedData{}
	})
	eh.RegisterEventData(Deleted, func() eh.EventData {
		return &DeletedData{}
	})
	eh.RegisterEventData(Updated, func() eh.EventData {
		return &UpdatedData{}
	})
}

// CreatedData is the event data for the Created event.
type CreatedData struct {
	Attributes PaymentAttributes `json:"attributes" bson:"attributes"`
}

// UpdatedData is the event data for the Updated event.
type UpdatedData struct {
	Attributes PaymentAttributes `json:"attributes" bson:"attributes"`
}

// DeletedData is the event data for the Deleted event
type DeletedData struct {
	Reason string `json:"reason" bson:"reason"`
}
