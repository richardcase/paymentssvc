package handlers

import (
	dom "github.com/richardcase/paymentssvc/pkg/domain"
	swag "github.com/richardcase/paymentssvc/pkg/gen/models"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
)

// DomainToSwaggerPayment transforms a payment from domain to swagger
func DomainToSwaggerPayment(domPayment *dom.Payment) *swag.Payment {
	idStr := strfmt.UUID(domPayment.ID.String())
	orgStr := strfmt.UUID(domPayment.OrganisationID.String())
	version := int32(domPayment.Version)

	return &swag.Payment{
		ID:             conv.UUID(idStr),
		OrganisationID: conv.UUID(orgStr),
		Version:        &version,
		Attributes:     DomainToSwaggerPaymentAttributes(domPayment.Attributes),
	}
}

// DomainToSwaggerPaymentAttributes transforms payment attributes from domain to swagger models
func DomainToSwaggerPaymentAttributes(domAttributes dom.PaymentAttributes) *swag.PaymentAttributes {
	return &swag.PaymentAttributes{
		Amount:           &domAttributes.Amount,
		Currency:         &domAttributes.Currency,
		PaymentScheme:    &domAttributes.PaymentScheme,
		PaymentType:      &domAttributes.PaymentType,
		BeneficiaryParty: DomainToSwaggerParty(domAttributes.BeneficiaryParty),
		DebtorParty:      DomainToSwaggerParty(domAttributes.DebtorParty),
	}
}

// DomainToSwaggerParty tranforms party from domain to swagger models
func DomainToSwaggerParty(domParty dom.Party) *swag.Party {
	return &swag.Party{
		AccountName:   domParty.AccountName,
		AccountNumber: domParty.AccountNumber,
	}
}

// SwaggerToDomainPaymentAttributes tranforms payment attributes from swagger to domain models
func SwaggerToDomainPaymentAttributes(swagAttributes *swag.PaymentAttributes) *dom.PaymentAttributes {
	return &dom.PaymentAttributes{
		Amount:           *swagAttributes.Amount,
		Currency:         *swagAttributes.Currency,
		PaymentScheme:    *swagAttributes.PaymentScheme,
		PaymentType:      *swagAttributes.PaymentType,
		BeneficiaryParty: SwaggerToDomainParty(swagAttributes.BeneficiaryParty),
		DebtorParty:      SwaggerToDomainParty(swagAttributes.DebtorParty),
	}
}

// SwaggerToDomainParty tranforms a part from swagger to domain model
func SwaggerToDomainParty(swagParty *swag.Party) dom.Party {
	return dom.Party{
		AccountName:   swagParty.AccountName,
		AccountNumber: swagParty.AccountNumber,
	}
}
