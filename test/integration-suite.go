package test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/cmd/di"
	healthcheck_ui "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/healthcheck/infra/ui"
	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"
	user_ui "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/infra/ui"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/test/helpers"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

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
	suite.FlushDb()
	suite.ArrangeWg = &sync.WaitGroup{}
}

func (suite *IntegrationSuite) TearDownTest() {
}

func (suite *IntegrationSuite) TearDownSuite() {
}

func (suite *IntegrationSuite) FlushDb() {
	suite.removeDbTable("users")
	suite.createUsersDbTable()
}

func (suite *IntegrationSuite) removeDbTable(table string) {
	_, err := suite.CommonServices.DynamoDbClient.DeleteTable(suite.Ctx, &dynamodb.DeleteTableInput{
		TableName: aws.String(table)})
	if err != nil {
		fmt.Printf("Couldn't delete table %v. Here's why: %v\n", table, err)
	}
}

func (suite *IntegrationSuite) createUsersDbTable() {
	var tableDesc *types.TableDescription
	table, err := suite.CommonServices.DynamoDbClient.CreateTable(suite.Ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("Id"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("Id"),
			KeyType:       types.KeyTypeHash,
		}},
		TableName: aws.String("users"),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	})
	if err != nil {
		fmt.Printf("Couldn't create table %v. Here's why: %v\n", "users", err)
	} else {
		waiter := dynamodb.NewTableExistsWaiter(suite.CommonServices.DynamoDbClient)
		err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
			TableName: aws.String("users")}, 5*time.Minute)
		if err != nil {
			fmt.Printf("Wait for table exists failed. Here's why: %v\n", err)
		}
		tableDesc = table.TableDescription
	}
	fmt.Println("DynamoDb table created: ", *tableDesc.TableName)
}

func (suite *IntegrationSuite) ExecuteJsonRequest(verb string, path string, body []byte, headers map[string]string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(verb, path, bytes.NewBuffer(body))
	if len(headers) != 0 {
		for headerName, value := range headers {
			req.Header.Set(headerName, value)
		}
	}

	assert.NoError(suite.T(), err)

	req.Header.Set("Content-Type", "application/json")
	return suite.ExecuteRequest(req)
}

func (suite *IntegrationSuite) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	if routerHandler == nil {
		routerHandler = suite.HttpServices.Router.GetMuxRouter()
	}

	routerHandler.ServeHTTP(rr, req)

	return rr
}

func (suite *IntegrationSuite) CheckResponse(expectedStatusCode int, expectedResponse string, response *httptest.ResponseRecorder, formats ...interface{}) {
	ja := jsonassert.New(suite.T())
	suite.CheckResponseCode(expectedStatusCode, response.Code)

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

func (suite *IntegrationSuite) CheckResponseCode(expected, actual int) {
	if expected != actual {
		suite.T().Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func (suite *IntegrationSuite) GivenUserWithId(userId string) {
	u := user_domain.NewUser(userId, helpers.RandomName(), time.Now())
	suite.CommonServices.Repositories.UserRepo.Save(suite.Ctx, *u)
}

func (suite *IntegrationSuite) CheckUserExists(userId string) {
	u, err := suite.CommonServices.Repositories.UserRepo.Find(suite.Ctx, userId)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), u)
}

func (suite *IntegrationSuite) CheckUserNotExists(userId string) {
	u, err := suite.CommonServices.Repositories.UserRepo.Find(suite.Ctx, userId)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), u)
}
