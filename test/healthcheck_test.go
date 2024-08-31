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
	response := suite.executeJsonRequest(
		http.MethodGet,
		"/healthcheck",
		nil,
		helpers.EmptyHeaders(),
	)

	suite.checkResponseCode(200, response.Code)
}
