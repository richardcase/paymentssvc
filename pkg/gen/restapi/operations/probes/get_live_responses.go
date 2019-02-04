// Code generated by go-swagger; DO NOT EDIT.

package probes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetLiveOKCode is the HTTP code returned for type GetLiveOK
const GetLiveOKCode int = 200

/*GetLiveOK service is alive

swagger:response getLiveOK
*/
type GetLiveOK struct {
}

// NewGetLiveOK creates GetLiveOK with default headers values
func NewGetLiveOK() *GetLiveOK {

	return &GetLiveOK{}
}

// WriteResponse to the client
func (o *GetLiveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetLiveBadRequestCode is the HTTP code returned for type GetLiveBadRequest
const GetLiveBadRequestCode int = 400

/*GetLiveBadRequest service isn't alive

swagger:response getLiveBadRequest
*/
type GetLiveBadRequest struct {
}

// NewGetLiveBadRequest creates GetLiveBadRequest with default headers values
func NewGetLiveBadRequest() *GetLiveBadRequest {

	return &GetLiveBadRequest{}
}

// WriteResponse to the client
func (o *GetLiveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
