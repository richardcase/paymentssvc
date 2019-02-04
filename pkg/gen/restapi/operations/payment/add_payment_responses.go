// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/richardcase/paymentssvc/pkg/gen/models"
)

// AddPaymentCreatedCode is the HTTP code returned for type AddPaymentCreated
const AddPaymentCreatedCode int = 201

/*AddPaymentCreated Payment has been created

swagger:response addPaymentCreated
*/
type AddPaymentCreated struct {
	/*The location to retrieve the payment instance

	 */
	Location string `json:"Location"`
}

// NewAddPaymentCreated creates AddPaymentCreated with default headers values
func NewAddPaymentCreated() *AddPaymentCreated {

	return &AddPaymentCreated{}
}

// WithLocation adds the location to the add payment created response
func (o *AddPaymentCreated) WithLocation(location string) *AddPaymentCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the add payment created response
func (o *AddPaymentCreated) SetLocation(location string) {
	o.Location = location
}

// WriteResponse to the client
func (o *AddPaymentCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AddPaymentBadRequestCode is the HTTP code returned for type AddPaymentBadRequest
const AddPaymentBadRequestCode int = 400

/*AddPaymentBadRequest Invalid input

swagger:response addPaymentBadRequest
*/
type AddPaymentBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddPaymentBadRequest creates AddPaymentBadRequest with default headers values
func NewAddPaymentBadRequest() *AddPaymentBadRequest {

	return &AddPaymentBadRequest{}
}

// WithPayload adds the payload to the add payment bad request response
func (o *AddPaymentBadRequest) WithPayload(payload *models.Error) *AddPaymentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add payment bad request response
func (o *AddPaymentBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddPaymentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddPaymentInternalServerErrorCode is the HTTP code returned for type AddPaymentInternalServerError
const AddPaymentInternalServerErrorCode int = 500

/*AddPaymentInternalServerError Internal Server Error

swagger:response addPaymentInternalServerError
*/
type AddPaymentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddPaymentInternalServerError creates AddPaymentInternalServerError with default headers values
func NewAddPaymentInternalServerError() *AddPaymentInternalServerError {

	return &AddPaymentInternalServerError{}
}

// WithPayload adds the payload to the add payment internal server error response
func (o *AddPaymentInternalServerError) WithPayload(payload *models.Error) *AddPaymentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add payment internal server error response
func (o *AddPaymentInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddPaymentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}