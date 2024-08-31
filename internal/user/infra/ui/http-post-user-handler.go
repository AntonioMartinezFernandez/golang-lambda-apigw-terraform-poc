package user_ui

import (
	"encoding/json"
	"io"
	"net/http"

	user_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/application"
	user_infra "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/infra"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	http_middlewares "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/http/middleware"
	json_schema "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/json-schema"

	"github.com/google/jsonapi"
)

const jsonSchema = "create-user.schema.json"

func NewPostUserHandler(
	commandBus bus.CommandBus,
	responseMiddleware *http_middlewares.JsonApiResponseMiddleware,
	jsonSchemaValidator json_schema.JsonSchemaValidator,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Extract request body
		body, err := io.ReadAll(r.Body)

		// Validate with JSON schema
		jsr, jsrErr := jsonSchemaValidator.Validate(body, jsonSchema)
		if !jsr.Valid() || jsrErr != nil {
			responseMiddleware.WriteErrorResponse(w, http_middlewares.BadRequestJsonApiHttpResponse("invalid body"), http.StatusBadRequest, err)
			return
		}

		// Unmarshall JSON to data struct
		var u user_infra.UserDto
		if err := json.Unmarshal(body, &u); err != nil {
			responseMiddleware.WriteErrorResponse(w, http_middlewares.BadRequestJsonApiHttpResponse("invalid body"), http.StatusBadRequest, err)
			return
		}

		// Create command and publish to command bus
		cmd := user_application.NewSaveUserCommand(u.Id, u.Name, u.Birthdate)
		cbErr := commandBus.Send(cmd)
		if cbErr != nil {
			responseMiddleware.WriteErrorResponse(w, []*jsonapi.ErrorObject{}, http.StatusInternalServerError, err)
			return
		}

		// Http response
		responseMiddleware.WriteResponse(w, nil, http.StatusCreated)
	}
}
