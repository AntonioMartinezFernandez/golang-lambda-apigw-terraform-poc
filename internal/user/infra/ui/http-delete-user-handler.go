package user_ui

import (
	"net/http"

	user_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/application"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	http_middlewares "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/http/middleware"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

func NewDeleteUserHandler(
	commandBus bus.CommandBus,
	responseMiddleware *http_middlewares.JsonApiResponseMiddleware,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["id"]

		command := user_application.NewDeleteUserCommand(userId)
		err := commandBus.Send(command)
		if err != nil {
			responseMiddleware.WriteErrorResponse(w, []*jsonapi.ErrorObject{}, http.StatusInternalServerError, err)
			return
		}

		responseMiddleware.WriteResponse(w, nil, http.StatusNoContent)
	}
}
