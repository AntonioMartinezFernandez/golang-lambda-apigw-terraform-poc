package http_middlewares

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/logger"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/google/jsonapi"
)

const (
	badRequestCode  = "bad_request"
	badRequestTitle = "Bad Request"

	conflictCode  = "conflict"
	conflictTitle = "Conflict"

	unauthorizedCode  = "unauthorized"
	unauthorizedTitle = "Unauthorized"

	forbiddenCode  = "forbidden"
	forbiddenTitle = "Forbidden"

	internalCode  = "internal_server_error"
	internalTitle = "Internal Server Error"

	notFoundCode  = "not_found"
	notFoundTitle = "Not Found"

	tooManyRequestsCode  = "too_many_requests"
	tooManyRequestsTitle = "Too Many Requests"
)

type JsonApiResponseMiddleware struct {
	logger *slog.Logger
}

func NewJsonApiResponseMiddleware(logger *slog.Logger) *JsonApiResponseMiddleware {
	return &JsonApiResponseMiddleware{logger: logger}
}

func (jar *JsonApiResponseMiddleware) WriteErrorResponse(w http.ResponseWriter, errors []*jsonapi.ErrorObject, httpStatus int, previousError error) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(httpStatus)
	jar.logError(previousError, w.Header().Get(HeaderXRequestID), httpStatus)
	if err := jsonapi.MarshalErrors(w, errors); err != nil {
		jar.logger.Error("unexpected error marshalling json api response error", "error", err, "correlation_id", w.Header().Get(HeaderXRequestID))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (jar *JsonApiResponseMiddleware) WriteResponse(w http.ResponseWriter, payload interface{}, httpStatus int) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(httpStatus)

	if payload == nil {
		return
	}

	if err := jsonapi.MarshalPayload(w, payload); err != nil {
		jar.logger.Error("unexpected error marshalling json api response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (jar *JsonApiResponseMiddleware) logError(err error, correlationId string, statusCode int) {
	if err == nil {
		return
	}

	if statusCode >= http.StatusInternalServerError {
		logger.LogErrors(logger.Error, err, jar.logger, map[string]interface{}{"correlation_id": correlationId})
	} else {
		logger.LogErrors(logger.Warning, err, jar.logger, map[string]interface{}{"correlation_id": correlationId})
	}
}

func BadRequestJsonApiHttpResponse(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   badRequestCode,
		Title:  badRequestTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusBadRequest),
	}}
}

func ConflictJsonApiHttpResponse(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   conflictCode,
		Title:  conflictTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusConflict),
	}}
}

func UnauthorizedRequestJsonApiHttpResponse(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   unauthorizedCode,
		Title:  unauthorizedTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusUnauthorized),
	}}
}

func ForbiddenRequestJsonApiHttpResponse(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   forbiddenCode,
		Title:  forbiddenTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusForbidden),
	}}
}

func InternalServerErrorJsonApiHttpResponse() []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   internalCode,
		Title:  internalTitle,
		Detail: internalTitle,
		Status: strconv.Itoa(http.StatusInternalServerError),
	}}
}

func NotFoundRequestJsonApiHttpResponse(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   notFoundCode,
		Title:  notFoundTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusNotFound),
	}}
}

func TooManyRequestsJsonApiHttpResponse(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   tooManyRequestsCode,
		Title:  tooManyRequestsTitle,
		Detail: detail,
		Status: strconv.Itoa(http.StatusTooManyRequests),
	}}
}
