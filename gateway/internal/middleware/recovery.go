package middleware

import (
	"net/http"

	"github.com/Shivam/distributed-api-gateway/internal/logger"
)

// RecoveryMiddleware prevents server crashes
func RecoveryMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {

				logger.Log.Error().
					Any("panic", err).
					Msg("panic recovered")

				http.Error(
					w,
					"Internal Server Error",
					http.StatusInternalServerError,
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}