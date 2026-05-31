package proxy

import (
	"net/http"
	"net/http/httputil"

	"github.com/Shivam/distributed-api-gateway/internal/loadbalancer"
)

// ProxyHandler handles load-balanced proxying
type ProxyHandler struct {
	LoadBalancer *loadbalancer.LoadBalancer
}

// NewProxy creates proxy handler
func NewProxy(targets []string) (*ProxyHandler, error) {

	lb, err := loadbalancer.NewLoadBalancer(targets)
	if err != nil {
		return nil, err
	}

	lb.StartHealthChecks()

	return &ProxyHandler{
		LoadBalancer: lb,
	}, nil
}

// ServeHTTP forwards request
func (p *ProxyHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {

	backend, err := p.LoadBalancer.NextBackend()
	if err != nil {

		http.Error(
			w,
			"No healthy backends available",
			http.StatusServiceUnavailable,
		)

		return
	}

	proxy := httputil.NewSingleHostReverseProxy(
		backend.URL,
	)

	// Called AFTER backend responds
	proxy.ModifyResponse = func(
		resp *http.Response,
	) error {

		// 5xx means backend failure
		if resp.StatusCode >= 500 {

			backend.RecordFailure()

		} else {

			backend.RecordSuccess()
		}

		return nil
	}

	// Called when proxy cannot reach backend
	proxy.ErrorHandler = func(
		w http.ResponseWriter,
		r *http.Request,
		err error,
	) {

		backend.RecordFailure()

		http.Error(
			w,
			"backend unavailable",
			http.StatusBadGateway,
		)
	}

	proxy.ServeHTTP(w, r)
}