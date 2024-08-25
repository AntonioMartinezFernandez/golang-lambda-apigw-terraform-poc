package http_errors

import (
	"net/http"
	"strconv"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/google/jsonapi"
)

func NewUnauthorized(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "unauthorized",
		Title:  "Unauthorized",
		Detail: detail,
		Status: strconv.Itoa(http.StatusUnauthorized),
	}}
}
