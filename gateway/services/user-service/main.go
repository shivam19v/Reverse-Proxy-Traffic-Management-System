package main

import (
	"encoding/json"
	"net/http"
)

// Response structure
type Response struct {
	Service string `json:"service"`
	Message string `json:"message"`
}

func main() {

	// Create route handler
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		response := Response{
			Service: "USER SERVICE",
			Message: "Users endpoint reached",
		}

		// Set response type as JSON
		w.Header().Set("Content-Type", "application/json")

		// Send JSON response
		json.NewEncoder(w).Encode(response)
	})

	// Start server on port 9001
	http.ListenAndServe(":9001", nil)
}