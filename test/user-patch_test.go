package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/test/helpers"

	"github.com/stretchr/testify/suite"
)

type UpdateUserSuite struct {
	IntegrationSuite
}

func (suite *UpdateUserSuite) SetupSuite() {
	suite.IntegrationSuite.SetupSuite()
}

func (suite *UpdateUserSuite) SetupTest() {
	suite.IntegrationSuite.SetupTest()
}

func TestUpdateUserSuite(t *testing.T) {
	suite.Run(t, new(UpdateUserSuite))
}

func (suite *UpdateUserSuite) TestHandlePatchUserRequest() {
	userId := "01J64V13D4AHZ61T4MD7Z53BVZ"
	suite.GivenUserWithId(userId)

	userName := helpers.RandomName()
	userBirthdate := helpers.RandomTimeRFC3339()

	requestBody := fmt.Sprintf(
		`{"id": "%s","name": "%s","birthdate": "%s"}`,
		userId,
		userName,
		userBirthdate,
	)

	response := suite.executeJsonRequest(
		http.MethodPatch,
		"/users",
		[]byte(requestBody),
		helpers.EmptyHeaders(),
	)

	suite.checkResponse(http.StatusNoContent, "", response)

	// TODO: check user saved in db
}
