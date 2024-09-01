package user_application

import (
	"context"
	"errors"
	"time"

	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
)

type UpdateUserCommandHandler struct {
	userRepo user_domain.UserRepository
}

func NewUpdateUserCommandHandler(userRepo user_domain.UserRepository) UpdateUserCommandHandler {
	return UpdateUserCommandHandler{
		userRepo: userRepo,
	}
}

func (h UpdateUserCommandHandler) Handle(c bus.Command) error {
	cmd, ok := c.(*UpdateUserCommand)
	if !ok {
		return bus.NewInvalidDto("invalid update user command")
	}

	ctx := context.Background()

	cmdData := cmd.Data()
	userId := cmdData["id"].(string)
	userName := cmdData["name"]
	userBirthdate := cmdData["birthdate"]

	user, err := h.userRepo.Find(ctx, userId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("error handling update user command because user not exists")
	}

	var name *string
	if !utils.InterfacePointerIsNil(userName) {
		name, _ = userName.(*string)
	}

	var birthdate *time.Time
	if !utils.InterfacePointerIsNil(userBirthdate) {
		bs, _ := userBirthdate.(*string)
		b, err := time.Parse(time.RFC3339, *bs)
		if err != nil {
			return err
		}
		birthdate = &b
	}

	user.Update(name, birthdate)

	return h.userRepo.Update(ctx, *user)
}
