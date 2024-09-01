package user_application_test

import (
	"context"
	"errors"
	"testing"
	"time"

	user_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/application"
	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"
	user_domain_mocks "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserCommandHandler(t *testing.T) {
	t.Run("update only user name", func(t *testing.T) {
		ctx := context.Background()
		originalBirthDate, _ := time.Parse(time.RFC3339, "1963-02-17T11:07:12Z")
		originalUser := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", "Lebron James", originalBirthDate)

		newName := "Michael Jordan"
		nameUpdatedUser := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", "Michael Jordan", originalBirthDate)

		repo := user_domain_mocks.NewUserRepository(t)
		repo.On("Find", ctx, "01J6J2VKXHR0A65AHG38J4RJB4").Return(originalUser, nil).Once()
		repo.On("Update", ctx, *nameUpdatedUser).Return(nil).Once()

		handler := user_application.NewUpdateUserCommandHandler(repo)

		cmd := user_application.NewUpdateUserCommand("01J6J2VKXHR0A65AHG38J4RJB4", &newName, nil)
		err := handler.Handle(cmd)
		assert.NoError(t, err)
	})

	t.Run("update only birthdate name", func(t *testing.T) {
		ctx := context.Background()
		originalBirthDate, _ := time.Parse(time.RFC3339, "1963-02-17T11:07:12Z")
		originalUser := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", "Lebron James", originalBirthDate)

		newBirthdate := "1984-12-30T17:04:05Z"
		parsedNewBirthdate, _ := time.Parse(time.RFC3339, newBirthdate)
		birthdateUpdatedUser := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", "Lebron James", parsedNewBirthdate)

		repo := user_domain_mocks.NewUserRepository(t)
		repo.On("Find", ctx, "01J6J2VKXHR0A65AHG38J4RJB4").Return(originalUser, nil).Once()
		repo.On("Update", ctx, *birthdateUpdatedUser).Return(nil).Once()

		handler := user_application.NewUpdateUserCommandHandler(repo)

		cmd := user_application.NewUpdateUserCommand("01J6J2VKXHR0A65AHG38J4RJB4", nil, &newBirthdate)
		err := handler.Handle(cmd)
		assert.NoError(t, err)
	})

	t.Run("update name and birthdate", func(t *testing.T) {
		ctx := context.Background()
		originalBirthDate, _ := time.Parse(time.RFC3339, "1963-02-17T11:07:12Z")
		originalUser := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", "Michael Jordan", originalBirthDate)

		newName := "Lebron James"
		newBirthdate := "1984-12-30T17:04:05Z"
		parsedNewBirthdate, _ := time.Parse(time.RFC3339, newBirthdate)
		updatedUser := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", newName, parsedNewBirthdate)

		repo := user_domain_mocks.NewUserRepository(t)
		repo.On("Find", ctx, "01J6J2VKXHR0A65AHG38J4RJB4").Return(originalUser, nil).Once()
		repo.On("Update", ctx, *updatedUser).Return(nil).Once()

		handler := user_application.NewUpdateUserCommandHandler(repo)

		cmd := user_application.NewUpdateUserCommand("01J6J2VKXHR0A65AHG38J4RJB4", &newName, &newBirthdate)
		err := handler.Handle(cmd)
		assert.NoError(t, err)
	})

	t.Run("users repository error on find", func(t *testing.T) {
		ctx := context.Background()

		newName := "Lebron James"
		newBirthdate := "1984-12-30T17:04:05Z"

		repo := user_domain_mocks.NewUserRepository(t)
		repo.On("Find", ctx, "01J6J2VKXHR0A65AHG38J4RJB4").Return(nil, errors.New("error finding")).Once()

		handler := user_application.NewUpdateUserCommandHandler(repo)

		cmd := user_application.NewUpdateUserCommand("01J6J2VKXHR0A65AHG38J4RJB4", &newName, &newBirthdate)
		err := handler.Handle(cmd)
		assert.Error(t, err)
	})

	t.Run("users repository error on update", func(t *testing.T) {
		ctx := context.Background()
		originalBirthDate, _ := time.Parse(time.RFC3339, "1963-02-17T11:07:12Z")
		originalUser := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", "Michael Jordan", originalBirthDate)

		newName := "Lebron James"
		newBirthdate := "1984-12-30T17:04:05Z"
		parsedNewBirthdate, _ := time.Parse(time.RFC3339, newBirthdate)
		updatedUser := user_domain.NewUser("01J6J2VKXHR0A65AHG38J4RJB4", newName, parsedNewBirthdate)

		repo := user_domain_mocks.NewUserRepository(t)
		repo.On("Find", ctx, "01J6J2VKXHR0A65AHG38J4RJB4").Return(originalUser, nil).Once()
		repo.On("Update", ctx, *updatedUser).Return(errors.New("error updating")).Once()

		handler := user_application.NewUpdateUserCommandHandler(repo)

		cmd := user_application.NewUpdateUserCommand("01J6J2VKXHR0A65AHG38J4RJB4", &newName, &newBirthdate)
		err := handler.Handle(cmd)
		assert.Error(t, err)
	})
}
