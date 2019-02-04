package handlers

import (
	"fmt"

	eh "github.com/looplab/eventhorizon"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	dom "github.com/richardcase/paymentssvc/pkg/domain"
	"github.com/richardcase/paymentssvc/pkg/gen/models"
	"github.com/richardcase/paymentssvc/pkg/gen/restapi/operations/payment"
)

const (
	// ErrFindingPaymentMsg is the error message when there was an issue
	// finding a payment
	ErrFindingPaymentMsg = "Error occured finding payment"
)

// AddPayment is the hnadler for adding a new payment
func (h *Handlers) AddPayment(params payment.AddPaymentParams) middleware.Responder {
	domAttr := SwaggerToDomainPaymentAttributes(params.Body)

	create := &dom.Create{
		Attributes: *domAttr,
		ID:         uuid.New(),
	}

	err := h.CommandHandler.HandleCommand(params.HTTPRequest.Context(), create)
	if err != nil {
		message := fmt.Sprintf("error creating new payment %s", err)
		h.Logger.WithError(err).Error("error creating new payment")
		return payment.NewAddPaymentInternalServerError().WithPayload(h.createError(message))
	}

	location := fmt.Sprintf("/payments/%s", create.ID.String())

	return payment.NewAddPaymentCreated().WithLocation(location)
}

// DeletePayment is the handler for deleting a payment
func (h *Handlers) DeletePayment(params payment.DeletePaymentParams) middleware.Responder {
	idStr := params.ID.String()

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.Logger.WithError(err).Error("error parsing payment id")
		return payment.NewDeletePaymentMethodNotAllowed().WithPayload(h.createError("error parsing payment id"))
	}

	_, err = h.Repo.Find(params.HTTPRequest.Context(), id)
	if err != nil {
		rrErr, ok := err.(eh.RepoError)
		if ok && rrErr.Err == eh.ErrEntityNotFound {
			errMessage := fmt.Sprintf("error getting payment with id %s", idStr)
			return payment.NewDeletePaymentNotFound().WithPayload(h.createError(errMessage))
		}

		h.Logger.WithError(err).Error(ErrFindingPaymentMsg)
		return payment.NewDeletePaymentInternalServerError().WithPayload(h.createError(ErrFindingPaymentMsg))
	}

	delete := &dom.Delete{
		Reason: "TODO: capture intent",
		ID:     id,
	}

	err = h.CommandHandler.HandleCommand(params.HTTPRequest.Context(), delete)
	if err != nil {
		message := fmt.Sprintf("error deleteing payment %s", err)
		h.Logger.WithError(err).Error("error updating payment")
		return payment.NewDeletePaymentInternalServerError().WithPayload(h.createError(message))
	}

	return payment.NewDeletePaymentNoContent()
}

// GetPaymentByID implemenets the handler to get payment by its id
func (h *Handlers) GetPaymentByID(params payment.GetPaymentByIDParams) middleware.Responder {
	idStr := params.ID.String()

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.Logger.WithError(err).Error("error parsing payment id")
		return payment.NewGetPaymentByIDBadRequest().WithPayload(h.createError("error parsing payment id"))
	}

	data, err := h.Repo.Find(params.HTTPRequest.Context(), id)
	if err != nil {
		rrErr, ok := err.(eh.RepoError)
		if ok && rrErr.Err == eh.ErrEntityNotFound {
			errMessage := fmt.Sprintf("error getting payment with id %s", idStr)
			return payment.NewGetPaymentByIDNotFound().WithPayload(h.createError(errMessage))
		}

		h.Logger.WithError(err).Error(ErrFindingPaymentMsg)
		return payment.NewGetPaymentByIDInternalServerError().WithPayload(h.createError(ErrFindingPaymentMsg))
	}

	domPayment := data.(*dom.Payment)
	converted := DomainToSwaggerPayment(domPayment)

	return payment.NewGetPaymentByIDOK().WithPayload(converted)
}

// ListPayments is the hnadler to return a list of payments
func (h *Handlers) ListPayments(params payment.ListPaymentsParams) middleware.Responder {
	data, err := h.Repo.FindAll(params.HTTPRequest.Context())
	if err != nil {
		errMesg := "error listing payments"
		h.Logger.WithError(err).Error(errMesg)
		return payment.NewListPaymentsInternalServerError().WithPayload(h.createError(errMesg))
	}

	payments := []*models.Payment{}

	if len(data) > 0 {
		for _, entity := range data {
			domPayment := entity.(*dom.Payment)
			converted := DomainToSwaggerPayment(domPayment)
			payments = append(payments, converted)
		}
	}

	return payment.NewListPaymentsOK().WithPayload(payments)
}

// UpdatePayment is the handler to update a payment
func (h *Handlers) UpdatePayment(params payment.UpdatePaymentParams) middleware.Responder {
	idStr := params.ID.String()

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.Logger.WithError(err).Error("error parsing payment id")
		return payment.NewUpdatePaymentBadRequest().WithPayload(h.createError("error parsing payment id"))
	}

	_, err = h.Repo.Find(params.HTTPRequest.Context(), id)
	if err != nil {
		rrErr, ok := err.(eh.RepoError)
		if ok && rrErr.Err == eh.ErrEntityNotFound {
			errMessage := fmt.Sprintf("error getting payment with id %s", idStr)
			return payment.NewUpdatePaymentNotFound().WithPayload(h.createError(errMessage))
		}

		h.Logger.WithError(err).Error(ErrFindingPaymentMsg)
		return payment.NewUpdatePaymentInternalServerError().WithPayload(h.createError(ErrFindingPaymentMsg))
	}

	domAttr := SwaggerToDomainPaymentAttributes(params.Body)

	update := &dom.Update{
		Attributes: *domAttr,
		ID:         id,
	}

	err = h.CommandHandler.HandleCommand(params.HTTPRequest.Context(), update)
	if err != nil {
		message := fmt.Sprintf("error updating payment %s", err)
		h.Logger.WithError(err).Error("error updating payment")
		return payment.NewUpdatePaymentInternalServerError().WithPayload(h.createError(message))
	}

	return payment.NewUpdatePaymentOK()
}
