package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (

	// Total requests
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gateway_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
)

// RegisterMetrics registers prometheus metrics
func RegisterMetrics() {

	prometheus.MustRegister(RequestsTotal)
}