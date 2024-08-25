package http_errors

import (
	"net/http"
	"strconv"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/google/jsonapi"
)

func NewForbidden(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "forbidden",
		Title:  "Forbidden",
		Detail: detail,
		Status: strconv.Itoa(http.StatusForbidden),
	}}
}
