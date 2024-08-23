// Package lambda defines top level Lambda handlers
package lambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	lambdaEchoAdapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

// EchoHandler is the specific lambda handler for echo
type EchoHandler struct {
	lambdaEchoAdapter *lambdaEchoAdapter.EchoLambda
}

// NewEchoHandler creates a new echo handler
func NewEchoHandler(lambdaEchoAdapter *lambdaEchoAdapter.EchoLambda) *EchoHandler {
	return &EchoHandler{
		lambdaEchoAdapter: lambdaEchoAdapter,
	}
}

// Handle proxies an API Gateway request
func (h *EchoHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return h.lambdaEchoAdapter.ProxyWithContext(ctx, request)
}
