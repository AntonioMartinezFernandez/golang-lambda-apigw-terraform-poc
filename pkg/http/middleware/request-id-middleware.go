package http_middlewares

import (
	"context"
	"net/http"

	"github.com/AntonioMartinezFernandez/golang-lambda-apigw-terraform-poc/pkg/utils"
)

const HeaderXRequestID = "X-Request-ID"

type IdentifierGenerator func() string

type RequestIdMiddleware struct {
	IdGenerator IdentifierGenerator
}

func NewRequestIdMiddleware(generator IdentifierGenerator) *RequestIdMiddleware {
	return &RequestIdMiddleware{IdGenerator: generator}
}

func (m *RequestIdMiddleware) RequestIdentifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IdGenerator == nil {
			m.IdGenerator = func() string {
				return utils.NewUuid().String()
			}
		}

		rid := r.Header.Get(HeaderXRequestID)

		if rid == "" {
			rid = m.IdGenerator()
		}

		r.Header.Set(HeaderXRequestID, rid)

		newRequest := r.WithContext(context.WithValue(r.Context(), "correlation_id", r.Header.Get("X-Request-ID")))

		w.Header().Set(HeaderXRequestID, rid)

		next.ServeHTTP(w, newRequest)
	})
}
