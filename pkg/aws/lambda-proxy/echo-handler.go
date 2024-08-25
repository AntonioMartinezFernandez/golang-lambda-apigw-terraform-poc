// Package lambda defines top level Lambda handlers
package lambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	aws_lambda_apiproxy_echo "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"github.com/labstack/echo/v4"
)

// EchoHandler is the specific lambda handler for echo
type EchoHandler struct {
	lambdaEchoAdapter *aws_lambda_apiproxy_echo.EchoLambda
}

// NewEchoHandler creates a new echo handler
func NewEchoHandler(e *echo.Echo) *EchoHandler {
	awsLambdaApiproxyEcho := aws_lambda_apiproxy_echo.New(e)

	return &EchoHandler{
		lambdaEchoAdapter: awsLambdaApiproxyEcho,
	}
}

// Handle proxies an API Gateway request
func (h *EchoHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return h.lambdaEchoAdapter.ProxyWithContext(ctx, request)
}
