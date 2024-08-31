package user_application_test

import (
	"errors"
	"testing"
	"time"

	user_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/application"
	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"
	user_domain_mocks "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain/mocks"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestSaveUserCommandHandler(t *testing.T) {
	birthDate, _ := time.Parse("2006-01-02 15:04:05", "1984-11-25 17:04:12")
	user := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", "John", birthDate)

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
		"should not return error when user is saved successfully": {
			expectations: func(
				repository *user_domain_mocks.UserRepository,
				user *user_domain.User,
				err error,
			) {
				repository.On("Save", *user).Return(nil).Once()
			},
			user:          user,
			expectedError: nil,
		},
		"should return error when repository fails to save": {
			expectations: func(
				repository *user_domain_mocks.UserRepository,
				user *user_domain.User,
				err error,
			) {
				repository.On("Save", *user).Return(err).Once()
			},
			user:          user,
			expectedError: errors.New("repository error"),
		},
	}

	for name, tst := range tests {
		t.Run(name, func(t *testing.T) {
			repo := user_domain_mocks.NewUserRepository(t)
			ulidProvider := utils.NewFixedUlidProvider("01J6J2VKXHR0A65AHG38J4RJB4")
			handler := user_application.NewSaveUserCommandHandler(repo, ulidProvider)

			tst.expectations(repo, tst.user, tst.expectedError)

			command := user_application.NewSaveUserCommand(user.Id(), user.Name(), user.Birthdate().Format("2006-01-02 15:04:05"))
			err := handler.Handle(command)

			if tst.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, err, tst.expectedError)
				return
			}

			assert.NoError(t, err)
		})
	}
}
