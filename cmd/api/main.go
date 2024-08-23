// Package main starts the example Go Lambda
package main

import (
	awsLambda "github.com/aws/aws-lambda-go/lambda"
	lambdaEchoAdapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/api"
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/internal/lambda"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	apiHandler := api.NewHandler()

	e.GET("/hello", apiHandler.Hello)
	e.POST("/say-my-name", apiHandler.SayMyName)

	lambdaEchoHandler := lambda.NewEchoHandler(lambdaEchoAdapter.New(e))
	awsLambda.Start(lambdaEchoHandler.Handle)
}
