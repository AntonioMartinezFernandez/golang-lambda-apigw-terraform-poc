package test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/cmd/di"
	healthcheck_ui "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/healthcheck/infra/ui"
	user_ui "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/infra/ui"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var routerHandler http.Handler

type IntegrationSuite struct {
	suite.Suite
	Ctx            context.Context
	CommonServices *di.CommonServices
	HttpServices   *di.HttpServices
	ArrangeWg      *sync.WaitGroup
}

func (suite *IntegrationSuite) SetupSuite() {
	if suite.CommonServices == nil {
		fmt.Println("CommonServices is nil")
		suite.CommonServices = di.InitWithEnvFile("../.env", "../.env.TEST")
	}
	if suite.HttpServices == nil {
		fmt.Println("HttpServices is nil")
		suite.HttpServices = di.InitHttpServices(suite.CommonServices)
	}
	suite.Ctx = context.Background()

	healthcheck_ui.RegisterHealthcheckRoutes(suite.HttpServices, suite.CommonServices)
	user_ui.RegisterUserRoutes(suite.HttpServices, suite.CommonServices)
}

func (suite *IntegrationSuite) SetupTest() {
	suite.ArrangeWg = &sync.WaitGroup{}
}

func (suite *IntegrationSuite) TearDownTest() {
}

func (suite *IntegrationSuite) TearDownSuite() {
}

func (suite *IntegrationSuite) executeJsonRequest(verb string, path string, body []byte, headers map[string]string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(verb, path, bytes.NewBuffer(body))
	if len(headers) != 0 {
		for headerName, value := range headers {
			req.Header.Set(headerName, value)
		}
	}

	assert.NoError(suite.T(), err)

	req.Header.Set("Content-Type", "application/json")
	return suite.executeRequest(req)
}

func (suite *IntegrationSuite) executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	if routerHandler == nil {
		routerHandler = suite.HttpServices.Router.GetMuxRouter()
	}

	routerHandler.ServeHTTP(rr, req)

	return rr
}

func (suite *IntegrationSuite) checkResponse(expectedStatusCode int, expectedResponse string, response *httptest.ResponseRecorder, formats ...interface{}) {
	ja := jsonassert.New(suite.T())
	suite.checkResponseCode(expectedStatusCode, response.Code)

	receivedResponse := response.Body.String()
	if receivedResponse == "" {
		assert.Equal(suite.T(), expectedResponse, receivedResponse)
		return
	}
	if formats != nil {
		ja.Assertf(receivedResponse, expectedResponse, formats)
	} else {
		ja.Assertf(receivedResponse, expectedResponse)
	}
}

func (suite *IntegrationSuite) checkResponseCode(expected, actual int) {
	if expected != actual {
		suite.T().Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
