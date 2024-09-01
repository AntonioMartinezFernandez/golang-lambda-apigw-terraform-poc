package test

import (
	"net/http"
	"testing"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/test/helpers"

	"github.com/stretchr/testify/suite"
)

type GetUserSuite struct {
	IntegrationSuite
}

func (suite *GetUserSuite) SetupSuite() {
	suite.IntegrationSuite.SetupSuite()
}

func (suite *GetUserSuite) SetupTest() {
	suite.IntegrationSuite.SetupTest()
}

func TestGetUserSuite(t *testing.T) {
	suite.Run(t, new(GetUserSuite))
}

func (suite *GetUserSuite) TestHandleGetUserRequest() {
	userId := "01J64V13D4AHZ61T4MD7Z53BVZ"
	suite.GivenUserWithId(userId)

	// Make http request
	response := suite.ExecuteJsonRequest(
		http.MethodGet,
		"/users/01J64V13D4AHZ61T4MD7Z53BVZ",
		nil,
		helpers.EmptyHeaders(),
	)

	suite.CheckResponse(
		http.StatusOK,
		`{"data":{"type":"id","id":"01J64V13D4AHZ61T4MD7Z53BVZ","attributes":{"birthdate":"<<PRESENCE>>","id":"01J64V13D4AHZ61T4MD7Z53BVZ","name":"<<PRESENCE>>"}}}`,
		response,
	)
}
