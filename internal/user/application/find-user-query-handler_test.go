package user_application_test

import (
	"context"
	"errors"
	"testing"
	"time"

	user_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/application"
	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"
	user_domain_mocks "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain/mocks"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestFindUserQueryHandler(t *testing.T) {
	birthDate, _ := time.Parse("2006-01-02 15:04:05", "1984-11-25 17:04:12")
	user := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", "John", birthDate)
	ctx := context.Background()

	assert.Equal(t, user.Id(), "01J6J2VKXHR0A65AHG38J4RJB4")
	assert.Equal(t, user.Name(), "John")
	assert.Equal(t, user.Birthdate(), birthDate)

	tests := map[string]struct {
		expectations func(
			repository *user_domain_mocks.UserRepository,
			user *user_domain.User,
			err error,
		)
		user          *user_domain.User
		expectedError error
	}{
		"should return user when user is found": {
			expectations: func(
				repository *user_domain_mocks.UserRepository,
				user *user_domain.User,
				err error,
			) {
				repository.On("Find", ctx, "01J6J2VKXHR0A65AHG38J4RJB4").Return(user, nil).Once()
			},
			user:          user,
			expectedError: nil,
		},
		"should return error when repository fails to find": {
			expectations: func(
				repository *user_domain_mocks.UserRepository,
				user *user_domain.User,
				err error,
			) {
				repository.On("Find", ctx, "01J6J2VKXHR0A65AHG38J4RJB4").Return(nil, err).Once()
			},
			user:          nil,
			expectedError: errors.New("repository error"),
		},
	}

	for name, tst := range tests {
		t.Run(name, func(t *testing.T) {
			repo := user_domain_mocks.NewUserRepository(t)
			ulidProvider := utils.NewFixedUlidProvider("01J6J2VKXHR0A65AHG38J4RJB4")
			handler := user_application.NewFindUserQueryHandler(repo, ulidProvider)

			tst.expectations(repo, tst.user, tst.expectedError)

			query := user_application.NewFindUserQuery(user.Id())
			queryResponse, err := handler.Handle(query)

			if tst.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, err, tst.expectedError)
				return
			}

			castedQueryResponse := queryResponse.(user_domain.User)
			assert.Equal(t, castedQueryResponse.Id(), tst.user.Id())
			assert.Equal(t, castedQueryResponse.Name(), tst.user.Name())
			assert.Equal(t, castedQueryResponse.Birthdate(), tst.user.Birthdate())

			assert.NoError(t, err)
		})
	}
}
