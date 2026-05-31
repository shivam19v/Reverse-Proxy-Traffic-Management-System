package loadbalancer

import (
	"net/http"
	"time"

	"github.com/Shivam/distributed-api-gateway/internal/logger"
)

// HealthCheckPeriod defines interval
const HealthCheckPeriod = 10 * time.Second

// StartHealthChecks starts backend monitoring
func (lb *LoadBalancer) StartHealthChecks() {

	ticker := time.NewTicker(
		HealthCheckPeriod,
	)

	// Run forever in background goroutine
	go func() {

		for range ticker.C {

			lb.CheckBackends()
		}
	}()
}

// CheckBackends checks all backends
func (lb *LoadBalancer) CheckBackends() {

	for _, backend := range lb.Backends {

		go lb.CheckBackend(backend)
	}
}

// CheckBackend checks single backend
func (lb *LoadBalancer) CheckBackend(
	backend *Backend,
) {

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	healthURL := backend.URL.String() + "/health"

	resp, err := client.Get(healthURL)

	alive := err == nil &&
		resp.StatusCode == http.StatusOK

	backend.SetAlive(alive)

	if alive {

		logger.Log.Info().
			Str("backend", backend.URL.String()).
			Msg("backend healthy")

	} else {

		logger.Log.Warn().
			Str("backend", backend.URL.String()).
			Msg("backend unhealthy")
	}
}