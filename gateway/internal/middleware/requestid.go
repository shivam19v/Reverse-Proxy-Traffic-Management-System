package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

// RequestIDMiddleware adds unique request ID
func RequestIDMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Generate unique request ID
		requestID := uuid.New().String()

		// Add request ID to request header
		r.Header.Set("X-Request-ID", requestID)

		// Add request ID to response header
		w.Header().Set("X-Request-ID", requestID)

		// Continue request flow
		next.ServeHTTP(w, r)
	})
}