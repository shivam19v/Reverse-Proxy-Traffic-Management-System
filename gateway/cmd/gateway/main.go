package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shivam/distributed-api-gateway/internal/config"
	"github.com/Shivam/distributed-api-gateway/internal/logger"
	"github.com/Shivam/distributed-api-gateway/internal/metrics"
	"github.com/Shivam/distributed-api-gateway/internal/redis"
	"github.com/Shivam/distributed-api-gateway/internal/router"
)

func main() {

	// Load configuration
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {

		logger.Log.Fatal().
			Err(err).
			Msg("failed to load config")
	}

	// Initialize Redis
	redis.InitRedis()

	// Register Prometheus metrics
	metrics.RegisterMetrics()

	// Setup router
	r, err := router.SetupRouter(cfg)
	if err != nil {

		logger.Log.Fatal().
			Err(err).
			Msg("failed to setup router")
	}

	// Configure HTTP server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Run server in goroutine
	go func() {

		logger.Log.Info().
			Msg("API Gateway running on :8080")

		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {

			logger.Log.Fatal().
				Err(err).
				Msg("server failed")
		}
	}()

	// Wait for shutdown signal
	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-stop

	logger.Log.Info().
		Msg("shutting down gateway")

	// Graceful shutdown timeout
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	// Shutdown server
	err = server.Shutdown(ctx)
	if err != nil {

		logger.Log.Fatal().
			Err(err).
			Msg("graceful shutdown failed")
	}

	// Close Redis connection
	err = redis.Client.Close()
	if err != nil {

		logger.Log.Error().
			Err(err).
			Msg("failed to close redis connection")
	}

	logger.Log.Info().
		Msg("gateway stopped")
}