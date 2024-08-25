package di

import (
	"log/slog"

	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"
	user_infra "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/infra"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/config"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/logger"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
)

type Repositories struct {
	UserRepo user_domain.UserRepository
}

type CommonServices struct {
	Config       config.Config
	Logger       *slog.Logger
	UlidProvider utils.UlidProvider
	CommandBus   *bus.CommandBus
	QueryBus     *bus.QueryBus
	Repositories *Repositories
}

func Init() *CommonServices {
	config := initConfig()
	logger := logger.NewJsonLogger(config.LogLevel)
	repositories := initRepositories()
	ulidProvider := utils.NewRandomUlidProvider()
	commandBus := bus.NewCommandBus()
	queryBus := bus.NewQueryBus()

	RegisterBusHandlers(config, logger, repositories, ulidProvider, queryBus, commandBus)

	return &CommonServices{
		Config:       config,
		Logger:       logger,
		UlidProvider: ulidProvider,
		CommandBus:   commandBus,
		QueryBus:     queryBus,
		Repositories: repositories,
	}
}

func initConfig() config.Config {
	return config.LoadEnvConfig()
}

func initRepositories() *Repositories {
	return &Repositories{
		UserRepo: user_infra.NewInMemoryUserRepository(),
	}
}
