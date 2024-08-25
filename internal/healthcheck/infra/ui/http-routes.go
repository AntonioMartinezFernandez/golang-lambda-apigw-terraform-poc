package healthcheck_ui

import (
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/cmd/di"
)

func RegisterHealthcheckRoutes(httpServices *di.HttpServices, commonServices *di.CommonServices) {
	httpServices.Router.Get(
		"/healthcheck",
		httpServices.DefaultRouteMatching,
		NewHealthcheckHandler(*commonServices.QueryBus, httpServices.JsonApiResponseMiddleware),
	)
}
