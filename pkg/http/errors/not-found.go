package http_errors

import (
	"net/http"
	"strconv"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/google/jsonapi"
)

func NewNotFound(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "not_found",
		Title:  "Not Found",
		Detail: detail,
		Status: strconv.Itoa(http.StatusNotFound),
	}}
}
