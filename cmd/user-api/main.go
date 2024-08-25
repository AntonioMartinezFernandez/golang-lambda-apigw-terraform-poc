package main

import (
	"flag"
	"fmt"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/cmd/di"

	healthcheck_ui "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/healthcheck/infra/ui"
	user_ui "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/user/infra/ui"

	lambda_proxy "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/aws/lambda-proxy"

	aws_lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// CLI flags
	var devMode bool
	flag.BoolVar(&devMode, "dev", false, "Define if the application is in development mode")
	flag.Parse()

	// Init Dependency Injection
	commonServices := di.Init()
	httpServices := di.InitHttpServices(commonServices)
	commonServices.Logger.Info(
		fmt.Sprintf("%s running", commonServices.Config.AppServiceName),
		"version", commonServices.Config.AppVersion,
		"environment", commonServices.Config.AppEnv,
	)

	// Register http routes
	healthcheck_ui.RegisterHealthcheckRoutes(httpServices, commonServices)
	user_ui.RegisterUserRoutes(httpServices, commonServices)

	// Define the execution mode
	switch devMode {
	case true:
		// Start as independent service
		httpServerStarterError := httpServices.Router.
			ListenAndServe(fmt.Sprintf("%s:%v", "0.0.0.0", commonServices.Config.HttpPort))
		if httpServerStarterError != nil {
			commonServices.Logger.Error("error starting http server", "error", httpServerStarterError)
			panic("error starting http server")
		}
	case false:
		// Start as lambda function
		lambdaGorillaMuxHandler := lambda_proxy.NewGorillaMuxHandler(httpServices.Router.GetMuxRouter())
		aws_lambda.Start(lambdaGorillaMuxHandler.Handle)
	}
}
