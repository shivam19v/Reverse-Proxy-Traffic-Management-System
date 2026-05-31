package router

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents health status
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthHandler checks gateway health
func HealthHandler(w http.ResponseWriter, r *http.Request) {

	response := HealthResponse{
		Status: "OK",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}