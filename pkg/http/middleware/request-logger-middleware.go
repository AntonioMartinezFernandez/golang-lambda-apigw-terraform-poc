package http_middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	http_pkg "github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/http"
)

type RequestLoggerMiddleware struct {
	l *slog.Logger
}

func NewRequestLoggerMiddleware(l *slog.Logger) *RequestLoggerMiddleware {
	return &RequestLoggerMiddleware{l: l}
}

func (rl *RequestLoggerMiddleware) BasicRequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         200,
		}

		username := "-"
		logUrl := *r.URL
		if logUrl.User != nil {
			if name := logUrl.User.Username(); name != "" {
				username = name
			}
		}

		ipAddress := http_pkg.ClientIp(r)

		start := time.Now()
		defer func() {
			totalTime := time.Since(start).Milliseconds()
			rl.l.Debug(
				fmt.Sprintf("%s %s %d %s %s %d %s", r.Method, r.RequestURI, recorder.Status, time.Now().Format(time.RFC3339), ipAddress, totalTime, r.Referer()),
				"remote_addr_ip", ipAddress,
				"remote_user", username,
				"request_time_d", totalTime,
				"status", recorder.Status,
				"request", r.RequestURI,
				"request_method", r.Method,
				"http_referrer", r.Referer(),
				"http_user_agent", r.Header.Get("User-Agent"),
				"response_content_type", w.Header().Get("Content-Type"),
				"correlation_id", w.Header().Get(HeaderXRequestID),
				"message", fmt.Sprintf("%s %s %d %s %s %d %s", r.Method, r.RequestURI, recorder.Status, time.Now().Format(time.RFC3339), ipAddress, totalTime, r.Referer()),
			)
		}()

		next.ServeHTTP(recorder, r)
	})
}
