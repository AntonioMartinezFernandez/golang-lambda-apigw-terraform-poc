package http_errors

import (
	"net/http"
	"strconv"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/google/jsonapi"
)

func NewTooManyRequests(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "too_many_requests",
		Title:  "Too many requests",
		Detail: detail,
		Status: strconv.Itoa(http.StatusTooManyRequests),
	}}
}
