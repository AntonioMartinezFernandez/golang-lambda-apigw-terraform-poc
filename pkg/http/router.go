package http_pkg

import (
	"context"

	"github.com/gorilla/mux"

	"net/http"
	"strings"
	"time"
)

type Router struct {
	muxRouter *mux.Router
	server    *http.Server
}
type Options struct {
	WriteTimeout int
	ReadTimeout  int
	Middlewares  []Middleware
}

const HeaderVersionName = "Accept-version"

type RouteMatching func(r *http.Request) bool

type Middleware func(next http.Handler) http.Handler

func NewRouter(options Options) Router {
	readTimeout := 30 * time.Second
	writeTimeout := 30 * time.Second

	if options.ReadTimeout != 0 {
		readTimeout = time.Duration(options.ReadTimeout) * time.Second
	}

	if options.WriteTimeout != 0 {
		writeTimeout = time.Duration(options.WriteTimeout) * time.Second
	}

	m := muxRouter()
	m.Use(enableCORS)

	for _, middleware := range options.Middlewares {
		m.Use(mux.MiddlewareFunc(middleware))
	}

	h := &http.Server{
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	return Router{m, h}
}

func muxRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

func DefaultRouter(serverWriteTimeOut int, serverReadTimeout int, middlewares ...Middleware) *Router {
	options := Options{
		WriteTimeout: serverReadTimeout,
		ReadTimeout:  serverWriteTimeOut,
		Middlewares:  middlewares,
	}

	rt := NewRouter(options)

	return &rt
}

func (router *Router) ListenAndServe(addr string) error {
	router.server.Addr = addr
	router.server.Handler = router.muxRouter

	// Use default options
	return router.server.ListenAndServe()
}

func (router *Router) Shutdown(ctx context.Context) error {
	return router.server.Shutdown(ctx)
}

func (router *Router) GetMuxRouter() *mux.Router {
	return router.muxRouter
}

// Add global middlewares
func (router *Router) AddMiddleware(middlewares ...Middleware) {
	for _, middleware := range middlewares {
		router.muxRouter.Use(mux.MiddlewareFunc(middleware))
	}
}

func (router *Router) Handle(method string, path string, handler http.Handler, routeMatching RouteMatching, middlewares ...Middleware) {
	router.handleMultipleMethods([]string{method}, path, handler, routeMatching, middlewares...)
}

func (router *Router) handleMultipleMethods(methods []string, path string, handler http.Handler, routeMatching RouteMatching, middlewares ...Middleware) {
	route := router.muxRouter.NewRoute()
	for _, middleware := range middlewares {
		route.Subrouter().Use(mux.MiddlewareFunc(middleware))
	}
	route.Handler(handler)

	if nil != routeMatching {
		route.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
			return routeMatching(r)
		})
	}

	if strings.HasSuffix(path, "*") {
		route.PathPrefix(path[:len(path)-1])
	} else {
		route.Path(path)
	}

	if len(methods) != 0 {
		route.Methods(methods...)
	}
}

func (router *Router) Get(path string, routeMatching RouteMatching, handler http.Handler, middlewares ...Middleware) {
	router.Handle("GET", path, handler, routeMatching, middlewares...)
}

func (router *Router) Post(path string, routeMatching RouteMatching, handler http.Handler, middlewares ...Middleware) {
	router.Handle("POST", path, handler, routeMatching, middlewares...)
}

func (router *Router) Put(path string, routeMatching RouteMatching, handler http.Handler, middlewares ...Middleware) {
	router.Handle("PUT", path, handler, routeMatching, middlewares...)
}

func (router *Router) Patch(path string, routeMatching RouteMatching, handler http.Handler, middlewares ...Middleware) {
	router.Handle("PATCH", path, handler, routeMatching, middlewares...)
}

func (router *Router) Delete(path string, routeMatching RouteMatching, handler http.Handler, middlewares ...Middleware) {
	router.Handle("DELETE", path, handler, routeMatching, middlewares...)
}

func (router *Router) Head(path string, routeMatching RouteMatching, handler http.Handler, middlewares ...Middleware) {
	router.Handle("HEAD", path, handler, routeMatching, middlewares...)
}

func (router *Router) Options(path string, routeMatching RouteMatching, handler http.Handler, middlewares ...Middleware) {
	router.Handle("OPTIONS", path, handler, routeMatching, middlewares...)
}

func (router *Router) Route(methods []string, path string, routeMatching RouteMatching, handler http.Handler, middlewares ...Middleware) {
	router.handleMultipleMethods(methods, path, handler, routeMatching, middlewares...)
}

func NewDefaultRouteMatching() RouteMatching {
	return NewHeaderRouteMatching(HeaderVersionName, "")
}

func NewHeaderRouteMatching(headerName string, headerValue string) RouteMatching {
	return func(r *http.Request) bool {
		return r.Header.Get(headerName) == headerValue
	}
}

func NewHeaderVersionMatching(headerValue string) RouteMatching {
	return NewHeaderRouteMatching(HeaderVersionName, headerValue)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow specified HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")

		// Allow specified headers
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// Allow credentials
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}
