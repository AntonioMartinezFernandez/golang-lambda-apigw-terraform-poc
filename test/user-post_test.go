package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/test/helpers"

	"github.com/stretchr/testify/suite"
)

type SaveUserSuite struct {
	IntegrationSuite
}

func (suite *SaveUserSuite) SetupSuite() {
	suite.IntegrationSuite.SetupSuite()
}

func (suite *SaveUserSuite) SetupTest() {
	suite.IntegrationSuite.SetupTest()
}

func TestSaveUserSuite(t *testing.T) {
	suite.Run(t, new(SaveUserSuite))
}

func (suite *SaveUserSuite) TestHandlePostUserRequest() {
	userId := utils.NewUlid().String()
	userName := helpers.RandomName()
	userBirthdate := helpers.RandomTimeRFC3339()

	requestBody := fmt.Sprintf(
		`{"id": "%s","name": "%s","birthdate": "%s"}`,
		userId,
		userName,
		userBirthdate,
	)

	response := suite.executeJsonRequest(
		http.MethodPost,
		"/user",
		[]byte(requestBody),
		helpers.EmptyHeaders(),
	)

	suite.checkResponse(http.StatusCreated, "", response)

	// TODO: check user saved in db
}
