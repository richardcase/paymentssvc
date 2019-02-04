// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/richardcase/paymentssvc/pkg/gen/models"
)

// ListPaymentsOKCode is the HTTP code returned for type ListPaymentsOK
const ListPaymentsOKCode int = 200

/*ListPaymentsOK successful operation

swagger:response listPaymentsOK
*/
type ListPaymentsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Payment `json:"body,omitempty"`
}

// NewListPaymentsOK creates ListPaymentsOK with default headers values
func NewListPaymentsOK() *ListPaymentsOK {

	return &ListPaymentsOK{}
}

// WithPayload adds the payload to the list payments o k response
func (o *ListPaymentsOK) WithPayload(payload []*models.Payment) *ListPaymentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list payments o k response
func (o *ListPaymentsOK) SetPayload(payload []*models.Payment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPaymentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Payment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// ListPaymentsBadRequestCode is the HTTP code returned for type ListPaymentsBadRequest
const ListPaymentsBadRequestCode int = 400

/*ListPaymentsBadRequest Invalid status value

swagger:response listPaymentsBadRequest
*/
type ListPaymentsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListPaymentsBadRequest creates ListPaymentsBadRequest with default headers values
func NewListPaymentsBadRequest() *ListPaymentsBadRequest {

	return &ListPaymentsBadRequest{}
}

// WithPayload adds the payload to the list payments bad request response
func (o *ListPaymentsBadRequest) WithPayload(payload *models.Error) *ListPaymentsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list payments bad request response
func (o *ListPaymentsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPaymentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListPaymentsInternalServerErrorCode is the HTTP code returned for type ListPaymentsInternalServerError
const ListPaymentsInternalServerErrorCode int = 500

/*ListPaymentsInternalServerError Internal Server Error

swagger:response listPaymentsInternalServerError
*/
type ListPaymentsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListPaymentsInternalServerError creates ListPaymentsInternalServerError with default headers values
func NewListPaymentsInternalServerError() *ListPaymentsInternalServerError {

	return &ListPaymentsInternalServerError{}
}

// WithPayload adds the payload to the list payments internal server error response
func (o *ListPaymentsInternalServerError) WithPayload(payload *models.Error) *ListPaymentsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list payments internal server error response
func (o *ListPaymentsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPaymentsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}