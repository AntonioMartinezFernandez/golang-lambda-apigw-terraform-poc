// Package lambda defines top level Lambda handlers
package lambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	aws_lambda_apiproxy_gorillamux "github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"

	"github.com/gorilla/mux"
)

// GorillaMuxHandler is the specific lambda handler for Gorilla Mux
type GorillaMuxHandler struct {
	lambdaGorillaMuxAdapter *aws_lambda_apiproxy_gorillamux.GorillaMuxAdapter
}

// NewGorillaMuxHandler creates a new Gorilla Mux handler
func NewGorillaMuxHandler(r *mux.Router) *GorillaMuxHandler {
	awsLambdaApiproxyGorillaMux := aws_lambda_apiproxy_gorillamux.New(r)

	return &GorillaMuxHandler{
		lambdaGorillaMuxAdapter: awsLambdaApiproxyGorillaMux,
	}
}

// Handle proxies an API Gateway request
func (h *GorillaMuxHandler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r, err := h.lambdaGorillaMuxAdapter.ProxyWithContext(ctx, *core.NewSwitchableAPIGatewayRequestV1(&request))
	return *r.Version1(), err
}
