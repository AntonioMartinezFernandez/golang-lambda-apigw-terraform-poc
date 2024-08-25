package user_ui

import (
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/cmd/di"
)

func RegisterUserRoutes(httpServices *di.HttpServices, commonServices *di.CommonServices) {
	httpServices.Router.Get(
		"/user/{id}",
		httpServices.DefaultRouteMatching,
		NewGetUserHandler(*commonServices.QueryBus, httpServices.JsonApiResponseMiddleware),
	)

	httpServices.Router.Post(
		"/user",
		httpServices.DefaultRouteMatching,
		NewPostUserHandler(*commonServices.CommandBus, httpServices.JsonApiResponseMiddleware),
	)
}
