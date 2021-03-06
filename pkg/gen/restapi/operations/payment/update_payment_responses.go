// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/richardcase/paymentssvc/pkg/gen/models"
)

// UpdatePaymentOKCode is the HTTP code returned for type UpdatePaymentOK
const UpdatePaymentOKCode int = 200

/*UpdatePaymentOK Payment was updated successfully

swagger:response updatePaymentOK
*/
type UpdatePaymentOK struct {
}

// NewUpdatePaymentOK creates UpdatePaymentOK with default headers values
func NewUpdatePaymentOK() *UpdatePaymentOK {

	return &UpdatePaymentOK{}
}

// WriteResponse to the client
func (o *UpdatePaymentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdatePaymentBadRequestCode is the HTTP code returned for type UpdatePaymentBadRequest
const UpdatePaymentBadRequestCode int = 400

/*UpdatePaymentBadRequest Invalid ID supplied

swagger:response updatePaymentBadRequest
*/
type UpdatePaymentBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePaymentBadRequest creates UpdatePaymentBadRequest with default headers values
func NewUpdatePaymentBadRequest() *UpdatePaymentBadRequest {

	return &UpdatePaymentBadRequest{}
}

// WithPayload adds the payload to the update payment bad request response
func (o *UpdatePaymentBadRequest) WithPayload(payload *models.Error) *UpdatePaymentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment bad request response
func (o *UpdatePaymentBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePaymentNotFoundCode is the HTTP code returned for type UpdatePaymentNotFound
const UpdatePaymentNotFoundCode int = 404

/*UpdatePaymentNotFound Payment not found

swagger:response updatePaymentNotFound
*/
type UpdatePaymentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePaymentNotFound creates UpdatePaymentNotFound with default headers values
func NewUpdatePaymentNotFound() *UpdatePaymentNotFound {

	return &UpdatePaymentNotFound{}
}

// WithPayload adds the payload to the update payment not found response
func (o *UpdatePaymentNotFound) WithPayload(payload *models.Error) *UpdatePaymentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment not found response
func (o *UpdatePaymentNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePaymentMethodNotAllowedCode is the HTTP code returned for type UpdatePaymentMethodNotAllowed
const UpdatePaymentMethodNotAllowedCode int = 405

/*UpdatePaymentMethodNotAllowed Validation exception

swagger:response updatePaymentMethodNotAllowed
*/
type UpdatePaymentMethodNotAllowed struct {
}

// NewUpdatePaymentMethodNotAllowed creates UpdatePaymentMethodNotAllowed with default headers values
func NewUpdatePaymentMethodNotAllowed() *UpdatePaymentMethodNotAllowed {

	return &UpdatePaymentMethodNotAllowed{}
}

// WriteResponse to the client
func (o *UpdatePaymentMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

// UpdatePaymentInternalServerErrorCode is the HTTP code returned for type UpdatePaymentInternalServerError
const UpdatePaymentInternalServerErrorCode int = 500

/*UpdatePaymentInternalServerError Internal Server Error

swagger:response updatePaymentInternalServerError
*/
type UpdatePaymentInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdatePaymentInternalServerError creates UpdatePaymentInternalServerError with default headers values
func NewUpdatePaymentInternalServerError() *UpdatePaymentInternalServerError {

	return &UpdatePaymentInternalServerError{}
}

// WithPayload adds the payload to the update payment internal server error response
func (o *UpdatePaymentInternalServerError) WithPayload(payload *models.Error) *UpdatePaymentInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update payment internal server error response
func (o *UpdatePaymentInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePaymentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
