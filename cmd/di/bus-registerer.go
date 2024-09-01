package di

import (
	"log/slog"

	healthcheck_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/healthcheck/application"
	user_application "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/application"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/cmd/config"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
)

func RegisterBusHandlers(
	cnf config.Config,
	l *slog.Logger,
	repositories *Repositories,
	ulidProvider utils.UlidProvider,
	qb *bus.QueryBus,
	cb *bus.CommandBus,
) {

	// Register Queries
	qb.Register(&healthcheck_application.GetHealthcheckQuery{}, healthcheck_application.NewGetHealthcheckQueryHandler(cnf.AppServiceName, ulidProvider))
	qb.Register(&user_application.FindUserQuery{}, user_application.NewFindUserQueryHandler(repositories.UserRepo, ulidProvider))

	// Register Commands
	cb.Register(&user_application.SaveUserCommand{}, user_application.NewSaveUserCommandHandler(repositories.UserRepo, ulidProvider))
	cb.Register(&user_application.UpdateUserCommand{}, user_application.NewUpdateUserCommandHandler(repositories.UserRepo))
	cb.Register(&user_application.DeleteUserCommand{}, user_application.NewDeleteUserCommandHandler(repositories.UserRepo))
}
