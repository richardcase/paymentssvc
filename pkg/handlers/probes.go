package handlers

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/richardcase/paymentssvc/pkg/gen/restapi/operations/probes"
)

// GetLive is the handler for the liveness probe
func (h *Handlers) GetLive(params probes.GetLiveParams) middleware.Responder {
	return probes.NewGetLiveOK()
}

// GetReady is the handler for the readiness probe
func (h *Handlers) GetReady(params probes.GetReadyParams) middleware.Responder {
	h.readyMu.Lock()
	defer h.readyMu.Unlock()

	if h.ready {
		return probes.NewGetReadyOK()
	}

	return probes.NewGetReadyBadRequest()
}
