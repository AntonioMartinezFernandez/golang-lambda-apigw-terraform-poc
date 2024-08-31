package user_application

import (
	"context"

	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
)

type FindUserQueryHandler struct {
	userRepo     user_domain.UserRepository
	ulidProvider utils.UlidProvider
}

func NewFindUserQueryHandler(userRepo user_domain.UserRepository, ulidProvider utils.UlidProvider) FindUserQueryHandler {
	return FindUserQueryHandler{userRepo: userRepo, ulidProvider: ulidProvider}
}

func (guqh FindUserQueryHandler) Handle(query bus.Query) (interface{}, error) {
	q, ok := query.(*FindUserQuery)
	if !ok {
		return nil, bus.NewInvalidDto("Invalid query")
	}

	ctx := context.Background()

	user, err := guqh.userRepo.Find(ctx, q.userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	return *user, nil
}
