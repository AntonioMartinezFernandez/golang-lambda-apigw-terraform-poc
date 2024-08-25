package user_ui

import (
	"encoding/json"
	"net/http"

	user_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/application"

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

		query := user_application.NewGetUserQuery(userId)
		response, err := queryBus.Dispatch(query)
		if err != nil {
			responseMiddleware.WriteErrorResponse(w, []*jsonapi.ErrorObject{}, http.StatusInternalServerError, err)
			return
		}
		getUserResponse := response.(user_application.GetUserResponse)

		responseMiddleware.WriteResponse(w, &getUserResponse, http.StatusOK)
	}
}

func NewPostUserHandler(
	commandBus bus.CommandBus,
	responseMiddleware *http_middlewares.JsonApiResponseMiddleware,
) http.HandlerFunc {
	type User struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Birthdate string `json:"birthdate"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var u User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&u); err != nil {
			responseMiddleware.WriteErrorResponse(w, http_middlewares.BadRequestJsonApiHttpResponse("invalid body"), http.StatusBadRequest, err)
			return
		}
		defer r.Body.Close()

		cmd := user_application.NewSaveUserCommand(u.Id, u.Name, u.Birthdate)
		err := commandBus.Send(cmd)
		if err != nil {
			responseMiddleware.WriteErrorResponse(w, []*jsonapi.ErrorObject{}, http.StatusInternalServerError, err)
			return
		}

		responseMiddleware.WriteResponse(w, nil, http.StatusCreated)
	}
}
