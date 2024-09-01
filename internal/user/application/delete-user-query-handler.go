package user_application

import (
	"context"

	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
)

type DeleteUserCommandHandler struct {
	userRepo user_domain.UserRepository
}

func NewDeleteUserCommandHandler(userRepo user_domain.UserRepository) DeleteUserCommandHandler {
	return DeleteUserCommandHandler{userRepo: userRepo}
}

func (guqh DeleteUserCommandHandler) Handle(command bus.Command) error {
	cmd, ok := command.(*DeleteUserCommand)
	if !ok {
		return bus.NewInvalidDto("Invalid command")
	}

	ctx := context.Background()

	cmdData := cmd.Data()
	userId := cmdData["id"].(string)

	return guqh.userRepo.Delete(ctx, userId)
}
