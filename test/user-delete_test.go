package test

import (
	"net/http"
	"testing"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/test/helpers"

	"github.com/stretchr/testify/suite"
)

type DeleteUserSuite struct {
	IntegrationSuite
}

func (suite *DeleteUserSuite) SetupSuite() {
	suite.IntegrationSuite.SetupSuite()
}

func (suite *DeleteUserSuite) SetupTest() {
	suite.IntegrationSuite.SetupTest()
}

func TestDeleteUserSuite(t *testing.T) {
	suite.Run(t, new(DeleteUserSuite))
}

func (suite *DeleteUserSuite) TestHandleDeleteUserRequest() {
	userId := "01J64V13D4AHZ61T4MD7Z53BVZ"
	suite.GivenUserWithId(userId)

	// Make http request
	response := suite.executeJsonRequest(
		http.MethodDelete,
		"/users/01J64V13D4AHZ61T4MD7Z53BVZ",
		nil,
		helpers.EmptyHeaders(),
	)

	suite.checkResponseCode(http.StatusNoContent, response.Code)

	// TODO: check user not present in db
}
