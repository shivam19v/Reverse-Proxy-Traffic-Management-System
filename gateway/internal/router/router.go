package router

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Shivam/distributed-api-gateway/internal/config"
	"github.com/Shivam/distributed-api-gateway/internal/middleware"
	"github.com/Shivam/distributed-api-gateway/internal/proxy"
)

// SetupRouter creates all routes
func SetupRouter(cfg *config.Config) (*mux.Router, error) {

	// Create router
	router := mux.NewRouter()

	// Apply middleware
	router.Use(middleware.RecoveryMiddleware)
	router.Use(middleware.RequestIDMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RateLimiterMiddleware)
	router.Use(middleware.TimeoutMiddleware)

	// Health check endpoint
	router.HandleFunc("/health", HealthHandler).
		Methods("GET")

	// Prometheus metrics endpoint
	router.Handle(
		"/metrics",
		promhttp.Handler(),
	)

	// Configure proxy routes
	for _, route := range cfg.Routes {

		// Create reverse proxy
		proxyHandler, err := proxy.NewProxy(route.Targets)
		if err != nil {
			return nil, err
		}

		// Register route
		router.PathPrefix(route.Path).
			Handler(proxyHandler)
	}

	return router, nil
}