package test

import (
	"net/http"
	"testing"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/test/helpers"

	"github.com/stretchr/testify/suite"
)

type GetHealthcheckSuite struct {
	IntegrationSuite
}

func (suite *GetHealthcheckSuite) SetupSuite() {
	suite.IntegrationSuite.SetupSuite()
}

func (suite *GetHealthcheckSuite) SetupTest() {
	suite.IntegrationSuite.SetupTest()
}

func TestGetHealthcheckSuite(t *testing.T) {
	suite.Run(t, new(GetHealthcheckSuite))
}

func (suite *GetHealthcheckSuite) TestHandleGetHealthcheckRequest() {
	// Make http request
	response := suite.ExecuteJsonRequest(
		http.MethodGet,
		"/healthcheck",
		nil,
		helpers.EmptyHeaders(),
	)

	suite.CheckResponse(
		http.StatusOK,
		`{"data":{"type":"healthcheck","id":"<<PRESENCE>>","attributes":{"service":"rest_api_golang_lambda","status":"OK"}}}`,
		response,
	)
}
