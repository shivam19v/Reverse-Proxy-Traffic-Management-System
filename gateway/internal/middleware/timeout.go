package middleware

import (
	"net/http"
	"time"
)

// TimeoutMiddleware adds request timeout
func TimeoutMiddleware(next http.Handler) http.Handler {

	return http.TimeoutHandler(
		next,
		5*time.Second,
		"Request Timeout",
	)
}