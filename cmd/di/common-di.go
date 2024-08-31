package di

import (
	"log/slog"

	user_domain "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/domain"
	user_infra "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/infra"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/config"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/bus"
	json_schema "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/json-schema"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/logger"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"

	"github.com/joho/godotenv"
)

type Repositories struct {
	UserRepo user_domain.UserRepository
}

type CommonServices struct {
	Config              config.Config
	Logger              *slog.Logger
	JsonSchemaValidator json_schema.JsonSchemaValidator
	UlidProvider        utils.UlidProvider
	CommandBus          *bus.CommandBus
	QueryBus            *bus.QueryBus
	Repositories        *Repositories
}

func Init() *CommonServices {
	config := initConfig()
	logger := logger.NewJsonLogger(config.LogLevel)
	jsonSchemaValidator := json_schema.NewJsonSchemaValidator(config.JsonSchemaBasePath)
	repositories := initRepositories()
	ulidProvider := utils.NewRandomUlidProvider()
	commandBus := bus.NewCommandBus()
	queryBus := bus.NewQueryBus()

	RegisterBusHandlers(config, logger, repositories, ulidProvider, queryBus, commandBus)

	return &CommonServices{
		Config:              config,
		Logger:              logger,
		JsonSchemaValidator: jsonSchemaValidator,
		UlidProvider:        ulidProvider,
		CommandBus:          commandBus,
		QueryBus:            queryBus,
		Repositories:        repositories,
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

func InitWithEnvFile(envFiles ...string) *CommonServices {
	err := godotenv.Overload(envFiles...)
	if err != nil {
		panic(err)
	}

	return Init()
}
