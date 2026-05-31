package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Shivam/distributed-api-gateway/internal/logger"
	"github.com/Shivam/distributed-api-gateway/internal/redis"
)

// Rate limit configuration
const (
	RequestLimit = 100
	WindowSize   = 60 * time.Second
)

// RateLimiterMiddleware limits requests per IP
func RateLimiterMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Client IP address
		clientIP := r.RemoteAddr

		// Redis key
		key := fmt.Sprintf(
			"rate_limit:%s",
			clientIP,
		)

		// Increment request count atomically
		count, err := redis.Client.Incr(
			redis.Ctx,
			key,
		).Result()

		if err != nil {

			logger.Log.Error().
				Err(err).
				Msg("redis rate limit failed")

			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)

			return
		}

		// Set expiration ONLY for first request
		if count == 1 {

			err = redis.Client.Expire(
				redis.Ctx,
				key,
				WindowSize,
			).Err()

			if err != nil {

				logger.Log.Error().
					Err(err).
					Msg("failed to set redis expiration")
			}
		}

		// Add rate limit headers
		w.Header().Set(
			"X-RateLimit-Limit",
			strconv.Itoa(RequestLimit),
		)

		w.Header().Set(
			"X-RateLimit-Remaining",
			strconv.FormatInt(
				int64(RequestLimit)-count,
				10,
			),
		)

		// Reject if exceeded
		if count > RequestLimit {

			logger.Log.Warn().
				Str("client_ip", clientIP).
				Int64("request_count", count).
				Msg("rate limit exceeded")

			http.Error(
				w,
				"Too Many Requests",
				http.StatusTooManyRequests,
			)

			return
		}

		next.ServeHTTP(w, r)
	})
}