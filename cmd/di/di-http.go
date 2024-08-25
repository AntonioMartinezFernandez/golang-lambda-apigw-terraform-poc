package di

import (
	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/config"
	pkg_http "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/http"
	http_middlewares "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/http/middleware"
)

type HttpServices struct {
	Router                    *pkg_http.Router
	DefaultRouteMatching      pkg_http.RouteMatching
	JsonApiResponseMiddleware *http_middlewares.JsonApiResponseMiddleware
}

func InitHttpServices(commonServices *CommonServices) *HttpServices {
	responseMiddleware := http_middlewares.NewJsonApiResponseMiddleware(commonServices.Logger)

	return &HttpServices{
		Router:                    NewRouter(commonServices.Config, commonServices),
		DefaultRouteMatching:      pkg_http.NewDefaultRouteMatching(),
		JsonApiResponseMiddleware: responseMiddleware,
	}
}

func NewRouter(config config.Config, commonServices *CommonServices) *pkg_http.Router {
	return pkg_http.DefaultRouter(
		config.HttpWriteTimeout,
		config.HttpReadTimeout,
		http_middlewares.NewRequestIdMiddleware(nil).RequestIdentifier,
		http_middlewares.NewRequestLoggerMiddleware(commonServices.Logger).BasicRequestLoggerMiddleware,
		http_middlewares.NewRequestPanicMiddleware(commonServices.Logger).RequestPanicHandler,
	)
}
