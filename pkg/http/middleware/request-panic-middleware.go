package http_middlewares

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type RequestPanicMiddleware struct {
	l *slog.Logger
}

func NewRequestPanicMiddleware(l *slog.Logger) *RequestPanicMiddleware {
	return &RequestPanicMiddleware{l: l}
}

func (rp *RequestPanicMiddleware) RequestPanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				rp.l.Error("Unhandled Error", "error", err)
				jsonBody, _ := json.Marshal(map[string]interface{}{
					"errors": []map[string]string{
						{
							"title": "Internal Server Error",
						},
					},
				})
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				if _, err := w.Write(jsonBody); err != nil {
					rp.l.Error("Could not write the response of the panic error", "error", err)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
