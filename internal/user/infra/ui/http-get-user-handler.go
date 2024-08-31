package user_ui

import (
	"net/http"

	user_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/application"
	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"
	user_infra "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/infra"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	http_middlewares "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/http/middleware"

	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
)

func NewGetUserHandler(
	queryBus bus.QueryBus,
	responseMiddleware *http_middlewares.JsonApiResponseMiddleware,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId := vars["id"]

		query := user_application.NewFindUserQuery(userId)
		queryResponse, err := queryBus.Dispatch(query)
		if err != nil {
			responseMiddleware.WriteErrorResponse(w, []*jsonapi.ErrorObject{}, http.StatusInternalServerError, err)
			return
		}

		user, ok := queryResponse.(user_domain.User)
		if !ok {
			responseMiddleware.WriteErrorResponse(w, []*jsonapi.ErrorObject{}, http.StatusInternalServerError, err)
			return
		}

		getUserResponse := user_infra.NewGetUserResponse(
			user.Id(),
			user.Id(),
			user.Name(),
			user.Birthdate().Format("2006-01-02 15:04:05"),
		)
		responseMiddleware.WriteResponse(w, &getUserResponse, http.StatusOK)
	}
}
