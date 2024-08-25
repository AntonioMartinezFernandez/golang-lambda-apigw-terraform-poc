package healthcheck_ui

import (
	"net/http"

	healthcheck_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/healthcheck/application"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	http_middlewares "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/http/middleware"

	"github.com/google/jsonapi"
)

func NewHealthcheckHandler(
	queryBus bus.QueryBus,
	responseMiddleware *http_middlewares.JsonApiResponseMiddleware,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := healthcheck_application.NewGetHealthcheckQuery()
		response, err := queryBus.Dispatch(query)
		if err != nil {
			responseMiddleware.WriteErrorResponse(w, []*jsonapi.ErrorObject{}, http.StatusInternalServerError, err)
			return
		}
		healthcheckResponse := response.(healthcheck_application.GetHealthcheckResponse)

		responseMiddleware.WriteResponse(w, &healthcheckResponse, http.StatusOK)
	}
}
