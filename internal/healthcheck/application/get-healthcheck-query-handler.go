package healthcheck_application

import (
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
)

type GetHealthcheckQueryHandler struct {
	serviceName  string
	ulidProvider utils.UlidProvider
}

func NewGetHealthcheckQueryHandler(serviceName string, ulidProvider utils.UlidProvider) GetHealthcheckQueryHandler {
	return GetHealthcheckQueryHandler{serviceName: serviceName, ulidProvider: ulidProvider}
}

func (q GetHealthcheckQueryHandler) Handle(query bus.Query) (interface{}, error) {
	_, ok := query.(*GetHealthcheckQuery)
	if !ok {
		return nil, bus.NewInvalidDto("Invalid query")
	}

	return NewGetHealthcheckResponse(q.ulidProvider.New().String(), "OK", q.serviceName), nil
}
