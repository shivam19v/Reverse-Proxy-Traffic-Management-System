package middleware

import (
	"net/http"
	"time"

	"github.com/Shivam/distributed-api-gateway/internal/logger"
)

// LoggingMiddleware logs incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		requestID := r.Header.Get("X-Request-ID")

		logger.Log.Info().
			Str("request_id", requestID).
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Dur("duration", duration).
			Msg("incoming request")
	})
}