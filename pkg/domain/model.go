package domain

import (
	"time"

	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
)

// Payment represents a payment
type Payment struct {
	ID             uuid.UUID         `json:"id" bson:"_id"`
	OrganisationID uuid.UUID         `json:"organisation_id" bson:"organisation_id"`
	Version        int               `json:"version" bson:"version"`
	Attributes     PaymentAttributes `json:"attributes" json:"attributes"`
	IsDeleted      bool              `json:"deleted"    bson:"deleted"`
	CreatedAt      time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at" bson:"updated_at"`
}

var _ = eh.Entity(&Payment{})
var _ = eh.Versionable(&Payment{})

// EntityID is the identifier of this entity instance
func (p *Payment) EntityID() uuid.UUID {
	return p.ID
}

// AggregateVersion is the version of this entity instance
func (p *Payment) AggregateVersion() int {
	return p.Version
}

// PaymentAttributes represents the attributes of a payment
type PaymentAttributes struct {
	Amount           float32 `json:"amount" bson:"amount"`
	BeneficiaryParty Party   `json:"beneficiary_party" bson:"beneficiary_party"`
	Currency         string  `json:"currency" bson:"currency"`
	DebtorParty      Party   `json:"debtor_party" bson:"debtor_party"`
	PaymentScheme    string  `json:"payment_scheme" bson:"payment_scheme"`
	PaymentType      string  `json:"payment_type" bson:"payment_type"`
}

// Party represents a payment party
type Party struct {
	AccountName   string `json:"account_name,omitempty" bson:"account_name"`
	AccountNumber string `json:"account_number,omitempty" bson:"account_number"`
}
