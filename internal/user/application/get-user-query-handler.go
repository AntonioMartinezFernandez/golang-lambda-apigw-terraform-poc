package user_application

import (
	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
)

type GetUserQueryHandler struct {
	userRepo     user_domain.UserRepository
	ulidProvider utils.UlidProvider
}

func NewGetUserQueryHandler(userRepo user_domain.UserRepository, ulidProvider utils.UlidProvider) GetUserQueryHandler {
	return GetUserQueryHandler{userRepo: userRepo, ulidProvider: ulidProvider}
}

func (guqh GetUserQueryHandler) Handle(query bus.Query) (interface{}, error) {
	q, ok := query.(*GetUserQuery)
	if !ok {
		return nil, bus.NewInvalidDto("Invalid query")
	}

	user, err := guqh.userRepo.Find(q.userId)
	if err != nil {
		return nil, err
	}

	return NewGetUserResponse(
		guqh.ulidProvider.New().String(),
		user.Id(),
		user.Name(),
		user.Birthdate().String(),
	), nil
}
