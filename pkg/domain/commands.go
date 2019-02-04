package domain

import (
	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
)

const (
	// CreateCommand is the type for the Create payment command
	CreateCommand = eh.CommandType("payment:create")
	// UpdateCommand is the type for the Update payment command
	UpdateCommand = eh.CommandType("payment:update")
	// DeleteCommand is the type for the Delete payment command
	DeleteCommand = eh.CommandType("payment:delete")
)

var _ = eh.Command(&Create{})
var _ = eh.Command(&Update{})
var _ = eh.Command(&Delete{})

// Create represents the create payment command
type Create struct {
	ID         uuid.UUID         `json:"id"`
	Attributes PaymentAttributes `json:"attributes"`
}

// AggregateType returns the aggrgate type the command applies to
func (c *Create) AggregateType() eh.AggregateType { return AggregateType }

// AggregateID returns the ID of the aggregate
func (c *Create) AggregateID() uuid.UUID { return c.ID }

// CommandType returns the type of the command
func (c *Create) CommandType() eh.CommandType { return CreateCommand }

// Update represents the update payment command
type Update struct {
	ID         uuid.UUID         `json:"id"`
	Attributes PaymentAttributes `json:"attributes"`
}

// AggregateType returns the aggrgate type the command applies to
func (u *Update) AggregateType() eh.AggregateType { return AggregateType }

// AggregateID returns the ID of the aggregate
func (u *Update) AggregateID() uuid.UUID { return u.ID }

// CommandType returns the type of the command
func (u *Update) CommandType() eh.CommandType { return UpdateCommand }

// Delete represents the delete payment command
type Delete struct {
	ID     uuid.UUID `json:"id"`
	Reason string    `json:"amount"`
}

// AggregateType returns the aggrgate type the command applies to
func (d *Delete) AggregateType() eh.AggregateType { return AggregateType }

// AggregateID returns the ID of the aggregate
func (d *Delete) AggregateID() uuid.UUID { return d.ID }

// CommandType returns the type of the command
func (d *Delete) CommandType() eh.CommandType { return DeleteCommand }
