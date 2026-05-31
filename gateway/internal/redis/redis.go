package redis

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/Shivam/distributed-api-gateway/internal/logger"
)

var (
	// Global Redis context
	Ctx = context.Background()

	// Global Redis client
	Client *goredis.Client
)

// InitRedis initializes Redis connection
func InitRedis() {

	Client = goredis.NewClient(&goredis.Options{

		Addr:         "redis:6379",
		Password:     "",
		DB:           0,

		PoolSize:     10,
		MinIdleConns: 5,

		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// Verify Redis connection
	_, err := Client.Ping(Ctx).Result()
	if err != nil {

		logger.Log.Fatal().
			Err(err).
			Msg("failed to connect to redis")
	}

	logger.Log.Info().
		Msg("redis connected successfully")
}