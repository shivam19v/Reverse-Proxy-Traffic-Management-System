package health

import (
	"net/http"
	"time"
)

// CheckBackendHealth checks service availability
func CheckBackendHealth(target string) bool {

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(target + "/health")
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == 200
}