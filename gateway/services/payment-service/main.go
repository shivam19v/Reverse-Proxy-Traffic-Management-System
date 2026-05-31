package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Service string `json:"service"`
	Message string `json:"message"`
}

func main() {

	http.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {

		response := Response{
			Service: "PAYMENT SERVICE",
			Message: "Payments endpoint reached",
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":9002", nil)
}