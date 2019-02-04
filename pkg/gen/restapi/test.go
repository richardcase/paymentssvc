package restapi

import (
	"net/http"

	loads "github.com/go-openapi/loads"
	"github.com/richardcase/paymentssvc/pkg/gen/restapi/operations"
)

func GetAPI() (*operations.PaymentsAPI, error) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	api := operations.NewPaymentsAPI(swaggerSpec)
	return api, nil
}

func GetAPIHandler(api *operations.PaymentsAPI) (http.Handler, error) {
	h := configureAPI(api)
	err := api.Validate()
	if err != nil {
		return nil, err
	}
	return h, nil
}
