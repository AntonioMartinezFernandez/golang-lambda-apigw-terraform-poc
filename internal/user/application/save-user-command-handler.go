package user_application

import (
	"context"
	"time"

	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
)

type SaveUserCommandHandler struct {
	userRepo     user_domain.UserRepository
	ulidProvider utils.UlidProvider
}

func NewSaveUserCommandHandler(userRepo user_domain.UserRepository, ulidProvider utils.UlidProvider) SaveUserCommandHandler {
	return SaveUserCommandHandler{
		userRepo:     userRepo,
		ulidProvider: ulidProvider,
	}
}

func (h SaveUserCommandHandler) Handle(c bus.Command) error {
	cmd, ok := c.(*SaveUserCommand)
	if !ok {
		return bus.NewInvalidDto("invalid save user command")
	}

	ctx := context.Background()

	cmdData := cmd.Data()
	userId := cmdData["id"].(string)
	userName := cmdData["name"].(string)
	userBirthdate, err := time.Parse("2006-01-02 15:04:05", cmdData["birthdate"].(string))
	if err != nil {
		return err
	}

	user := user_domain.NewUser(userId, userName, userBirthdate)

	return h.userRepo.Save(ctx, *user)
}
