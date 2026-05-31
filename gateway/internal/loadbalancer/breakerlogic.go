package loadbalancer

import (
	"fmt"
	"time"
)

// IsAvailable checks if backend can receive traffic
func (b *Backend) IsAvailable() bool {

	if !b.IsAlive() {
		return false
	}

	breaker := b.Breaker

	breaker.Mutex.Lock()
	defer breaker.Mutex.Unlock()

	if breaker.State == StateClosed {
		return true
	}

	if breaker.State == StateOpen {

		// cooldown expired
		if time.Since(
			breaker.LastFailureTime,
		) > breaker.ResetTimeout {

			fmt.Println(
				"CIRCUIT BREAKER RESET:",
				b.URL.String(),
			)

			breaker.State = StateClosed
			breaker.Failures = 0

			return true
		}

		return false
	}

	return true
}

// RecordFailure records backend failure
func (b *Backend) RecordFailure() {

	breaker := b.Breaker

	breaker.Mutex.Lock()
	defer breaker.Mutex.Unlock()

	breaker.Failures++

	fmt.Println(
		"Backend failure:",
		b.URL.String(),
		"count:",
		breaker.Failures,
	)

	if breaker.Failures >= breaker.FailureThreshold {

		if breaker.State != StateOpen {

			fmt.Println(
				"CIRCUIT BREAKER OPENED:",
				b.URL.String(),
			)
		}

		breaker.State = StateOpen
		breaker.LastFailureTime = time.Now()
	}
}

// RecordSuccess resets failures
func (b *Backend) RecordSuccess() {

	breaker := b.Breaker

	breaker.Mutex.Lock()
	defer breaker.Mutex.Unlock()

	// only log if recovering from failures
	if breaker.Failures > 0 {

		fmt.Println(
			"Backend recovered:",
			b.URL.String(),
		)
	}

	breaker.Failures = 0
	breaker.State = StateClosed
}