package loadbalancer

import (
	"sync"
	"time"
)

// CircuitBreaker states
const (
	StateClosed = "CLOSED"
	StateOpen   = "OPEN"
)

// CircuitBreaker protects failing backends
type CircuitBreaker struct {

	// Current breaker state
	State string

	// Failure count
	Failures int

	// Max failures before opening
	FailureThreshold int

	// How long to stay open
	ResetTimeout time.Duration

	// When breaker opened
	LastFailureTime time.Time

	// Mutex for concurrent safety
	Mutex sync.Mutex
}

// NewCircuitBreaker creates breaker
func NewCircuitBreaker() *CircuitBreaker {

	return &CircuitBreaker{
		State:            StateClosed,
		FailureThreshold: 3,
		ResetTimeout:     15 * time.Second,
	}
}