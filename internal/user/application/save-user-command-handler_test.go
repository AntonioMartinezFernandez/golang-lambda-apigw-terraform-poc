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

func TestSaveUserCommandHandler(t *testing.T) {
	birthDate, _ := time.Parse(time.RFC3339, "1984-12-30T17:04:05Z")
	user := user_domain.NewUser("01J63630X372YYYR4CTFP1ZGGZ", "Lebron James", birthDate)
	ctx := context.Background()

	assert.Equal(t, user.Id(), "01J63630X372YYYR4CTFP1ZGGZ")
	assert.Equal(t, user.Name(), "Lebron James")
	assert.Equal(t, user.Birthdate().Format(time.RFC3339), birthDate.Format(time.RFC3339))

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
				repository.On("Save", ctx, *user).Return(nil).Once()
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
				repository.On("Save", ctx, *user).Return(err).Once()
			},
			user:          user,
			expectedError: errors.New("repository error"),
		},
	}

	for name, tst := range tests {
		t.Run(name, func(t *testing.T) {
			repo := user_domain_mocks.NewUserRepository(t)
			ulidProvider := utils.NewFixedUlidProvider("01J63630X372YYYR4CTFP1ZGGZ")
			handler := user_application.NewSaveUserCommandHandler(repo, ulidProvider)

			tst.expectations(repo, tst.user, tst.expectedError)

			command := user_application.NewSaveUserCommand(user.Id(), user.Name(), user.Birthdate().Format(time.RFC3339))
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
