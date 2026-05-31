package loadbalancer

import (
	"errors"
	"net/url"
	"sync"
	"sync/atomic"
)

// Backend represents a backend server
type Backend struct {

	// Backend URL
	URL *url.URL

	// Health status
	Alive bool

	// Protects Alive state
	Mutex sync.RWMutex

	// Circuit breaker
	Breaker *CircuitBreaker
}

// SetAlive safely updates backend health
func (b *Backend) SetAlive(alive bool) {

	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	b.Alive = alive
}

// IsAlive safely reads backend health
func (b *Backend) IsAlive() bool {

	b.Mutex.RLock()
	defer b.Mutex.RUnlock()

	return b.Alive
}

// LoadBalancer manages backend pool
type LoadBalancer struct {

	// All backend servers
	Backends []*Backend

	// Atomic round-robin counter
	Counter uint64
}

// NewLoadBalancer creates load balancer
func NewLoadBalancer(targets []string) (*LoadBalancer, error) {

	var backends []*Backend

	for _, target := range targets {

		parsedURL, err := url.Parse(target)
		if err != nil {
			return nil, err
		}

		backend := &Backend{
			URL:     parsedURL,
			Alive:   true,
			Breaker: NewCircuitBreaker(),
		}

		backends = append(backends, backend)
	}

	lb := &LoadBalancer{
		Backends: backends,
	}

	return lb, nil
}

// NextBackend returns next healthy backend
func (lb *LoadBalancer) NextBackend() (*Backend, error) {

	total := len(lb.Backends)

	if total == 0 {
		return nil, errors.New("no backends available")
	}

	// Try all backends
	for i := 0; i < total; i++ {

		// Atomic counter increment
		index := atomic.AddUint64(
			&lb.Counter,
			1,
		)

		backend := lb.Backends[
			int(index)%total,
		]

		// Return only healthy backends
		if backend.IsAvailable() {
			return backend, nil
		}
	}

	return nil, errors.New("no healthy backends available")
}

func (lb *LoadBalancer) NextBackendExcept(
	excluded *Backend,
) (*Backend, error) {

	total := len(lb.Backends)

	for i := 0; i < total; i++ {

		backend, err := lb.NextBackend()
		if err != nil {
			return nil, err
		}

		if backend != excluded {
			return backend, nil
		}
	}

	return nil, errors.New(
		"no alternative backend available",
	)
}