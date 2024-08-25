package http_errors

import (
	"net/http"
	"strconv"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/google/jsonapi"
)

func NewConflict(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "conflict",
		Title:  "Conflict",
		Detail: detail,
		Status: strconv.Itoa(http.StatusConflict),
	}}
}
