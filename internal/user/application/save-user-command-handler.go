package user_application

import (
	"time"

	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
)

type SaveUserCommandHandler struct {
	userRepo user_domain.UserRepository
}

func NewSaveUserCommandHandler(userRepo user_domain.UserRepository) SaveUserCommandHandler {
	return SaveUserCommandHandler{
		userRepo: userRepo,
	}
}

func (h SaveUserCommandHandler) Handle(c bus.Command) error {
	cmd, ok := c.(*SaveUserCommand)
	if !ok {
		return bus.NewInvalidDto("invalid save user command")
	}

	cmdData := cmd.Data()
	userId := cmdData["id"].(string)
	userName := cmdData["name"].(string)
	userBirthdate, err := time.Parse("2006-01-02 15:04:05", cmdData["birthdate"].(string))
	if err != nil {
		return err
	}

	user := user_domain.NewUser(userId, userName, userBirthdate)

	return h.userRepo.Save(*user)
}
