package user_ui

import (
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/cmd/di"
)

func RegisterUserRoutes(httpServices *di.HttpServices, commonServices *di.CommonServices) {
	httpServices.Router.Get(
		"/users/{id}",
		httpServices.DefaultRouteMatching,
		NewGetUserHandler(*commonServices.QueryBus, httpServices.JsonApiResponseMiddleware),
	)

	httpServices.Router.Post(
		"/users",
		httpServices.DefaultRouteMatching,
		NewPostUserHandler(
			*commonServices.CommandBus,
			httpServices.JsonApiResponseMiddleware,
			commonServices.JsonSchemaValidator,
		),
	)

	httpServices.Router.Patch(
		"/users",
		httpServices.DefaultRouteMatching,
		NewPatchUserHandler(
			*commonServices.CommandBus,
			httpServices.JsonApiResponseMiddleware,
			commonServices.JsonSchemaValidator,
		),
	)

	httpServices.Router.Delete(
		"/users/{id}",
		httpServices.DefaultRouteMatching,
		NewDeleteUserHandler(*commonServices.CommandBus, httpServices.JsonApiResponseMiddleware),
	)
}
