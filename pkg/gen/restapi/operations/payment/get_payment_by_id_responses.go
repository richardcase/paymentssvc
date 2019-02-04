// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/richardcase/paymentssvc/pkg/gen/models"
)

// GetPaymentByIDOKCode is the HTTP code returned for type GetPaymentByIDOK
const GetPaymentByIDOKCode int = 200

/*GetPaymentByIDOK payment was found

swagger:response getPaymentByIdOK
*/
type GetPaymentByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Payment `json:"body,omitempty"`
}

// NewGetPaymentByIDOK creates GetPaymentByIDOK with default headers values
func NewGetPaymentByIDOK() *GetPaymentByIDOK {

	return &GetPaymentByIDOK{}
}

// WithPayload adds the payload to the get payment by Id o k response
func (o *GetPaymentByIDOK) WithPayload(payload *models.Payment) *GetPaymentByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment by Id o k response
func (o *GetPaymentByIDOK) SetPayload(payload *models.Payment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentByIDBadRequestCode is the HTTP code returned for type GetPaymentByIDBadRequest
const GetPaymentByIDBadRequestCode int = 400

/*GetPaymentByIDBadRequest Invalid ID supplied

swagger:response getPaymentByIdBadRequest
*/
type GetPaymentByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPaymentByIDBadRequest creates GetPaymentByIDBadRequest with default headers values
func NewGetPaymentByIDBadRequest() *GetPaymentByIDBadRequest {

	return &GetPaymentByIDBadRequest{}
}

// WithPayload adds the payload to the get payment by Id bad request response
func (o *GetPaymentByIDBadRequest) WithPayload(payload *models.Error) *GetPaymentByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment by Id bad request response
func (o *GetPaymentByIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentByIDNotFoundCode is the HTTP code returned for type GetPaymentByIDNotFound
const GetPaymentByIDNotFoundCode int = 404

/*GetPaymentByIDNotFound Payment not found

swagger:response getPaymentByIdNotFound
*/
type GetPaymentByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPaymentByIDNotFound creates GetPaymentByIDNotFound with default headers values
func NewGetPaymentByIDNotFound() *GetPaymentByIDNotFound {

	return &GetPaymentByIDNotFound{}
}

// WithPayload adds the payload to the get payment by Id not found response
func (o *GetPaymentByIDNotFound) WithPayload(payload *models.Error) *GetPaymentByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment by Id not found response
func (o *GetPaymentByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentByIDInternalServerErrorCode is the HTTP code returned for type GetPaymentByIDInternalServerError
const GetPaymentByIDInternalServerErrorCode int = 500

/*GetPaymentByIDInternalServerError Internal Server Error

swagger:response getPaymentByIdInternalServerError
*/
type GetPaymentByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPaymentByIDInternalServerError creates GetPaymentByIDInternalServerError with default headers values
func NewGetPaymentByIDInternalServerError() *GetPaymentByIDInternalServerError {

	return &GetPaymentByIDInternalServerError{}
}

// WithPayload adds the payload to the get payment by Id internal server error response
func (o *GetPaymentByIDInternalServerError) WithPayload(payload *models.Error) *GetPaymentByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment by Id internal server error response
func (o *GetPaymentByIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}