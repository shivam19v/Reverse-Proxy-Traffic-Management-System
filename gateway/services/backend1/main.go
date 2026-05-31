package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("health check hit backend1")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("backend1 healthy"))
		},
	)

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("request hit backend1")

			fmt.Fprintf(
				w,
				"response from backend1\n",
			)
		},
	)

	fmt.Println("backend1 running on :9001")

	http.ListenAndServe(
		":9001",
		nil,
	)
}