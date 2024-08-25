package http_errors

import (
	"net/http"
	"strconv"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/google/jsonapi"
)

func NewInternalServerError() []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "internal_server_error",
		Title:  "Internal Server Error",
		Detail: "Internal Server Error",
		Status: strconv.Itoa(http.StatusInternalServerError),
	}}
}
