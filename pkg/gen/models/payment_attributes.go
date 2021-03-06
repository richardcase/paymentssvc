// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentAttributes payment attributes
// swagger:model PaymentAttributes
type PaymentAttributes struct {

	// amount
	// Required: true
	Amount *float32 `json:"amount"`

	// beneficiary party
	// Required: true
	BeneficiaryParty *Party `json:"beneficiary_party"`

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// debtor party
	// Required: true
	DebtorParty *Party `json:"debtor_party"`

	// Which payment scheme are we using
	// Required: true
	// Enum: [FPS BACS CHAPS SEPA]
	PaymentScheme *string `json:"payment_scheme"`

	// payment type
	// Required: true
	// Enum: [Credit Debit]
	PaymentType *string `json:"payment_type"`
}

// Validate validates this payment attributes
func (m *PaymentAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAttributes) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributes) validateBeneficiaryParty(formats strfmt.Registry) error {

	if err := validate.Required("beneficiary_party", "body", m.BeneficiaryParty); err != nil {
		return err
	}

	if m.BeneficiaryParty != nil {
		if err := m.BeneficiaryParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiary_party")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAttributes) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAttributes) validateDebtorParty(formats strfmt.Registry) error {

	if err := validate.Required("debtor_party", "body", m.DebtorParty); err != nil {
		return err
	}

	if m.DebtorParty != nil {
		if err := m.DebtorParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtor_party")
			}
			return err
		}
	}

	return nil
}

var paymentAttributesTypePaymentSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FPS","BACS","CHAPS","SEPA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentAttributesTypePaymentSchemePropEnum = append(paymentAttributesTypePaymentSchemePropEnum, v)
	}
}

const (

	// PaymentAttributesPaymentSchemeFPS captures enum value "FPS"
	PaymentAttributesPaymentSchemeFPS string = "FPS"

	// PaymentAttributesPaymentSchemeBACS captures enum value "BACS"
	PaymentAttributesPaymentSchemeBACS string = "BACS"

	// PaymentAttributesPaymentSchemeCHAPS captures enum value "CHAPS"
	PaymentAttributesPaymentSchemeCHAPS string = "CHAPS"

	// PaymentAttributesPaymentSchemeSEPA captures enum value "SEPA"
	PaymentAttributesPaymentSchemeSEPA string = "SEPA"
)

// prop value enum
func (m *PaymentAttributes) validatePaymentSchemeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentAttributesTypePaymentSchemePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentAttributes) validatePaymentScheme(formats strfmt.Registry) error {

	if err := validate.Required("payment_scheme", "body", m.PaymentScheme); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentSchemeEnum("payment_scheme", "body", *m.PaymentScheme); err != nil {
		return err
	}

	return nil
}

var paymentAttributesTypePaymentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Credit","Debit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentAttributesTypePaymentTypePropEnum = append(paymentAttributesTypePaymentTypePropEnum, v)
	}
}

const (

	// PaymentAttributesPaymentTypeCredit captures enum value "Credit"
	PaymentAttributesPaymentTypeCredit string = "Credit"

	// PaymentAttributesPaymentTypeDebit captures enum value "Debit"
	PaymentAttributesPaymentTypeDebit string = "Debit"
)

// prop value enum
func (m *PaymentAttributes) validatePaymentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentAttributesTypePaymentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentAttributes) validatePaymentType(formats strfmt.Registry) error {

	if err := validate.Required("payment_type", "body", m.PaymentType); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentTypeEnum("payment_type", "body", *m.PaymentType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
