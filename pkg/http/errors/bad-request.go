package http_errors

import (
	"net/http"
	"strconv"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/google/jsonapi"
)

func NewBadRequest(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "bad_request",
		Title:  "Bad Request",
		Detail: detail,
		Status: strconv.Itoa(http.StatusBadRequest),
	}}
}

func NewBadRequestCustom(code string, title string, detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   code,
		Title:  title,
		Detail: detail,
		Status: strconv.Itoa(http.StatusBadRequest),
	}}
}
