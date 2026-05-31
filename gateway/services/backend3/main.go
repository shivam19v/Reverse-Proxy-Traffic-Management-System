package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("health check hit backend3")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("backend3 healthy"))
		},
	)

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("request hit backend3")

			fmt.Fprintf(
				w,
				"response from backend3\n",
			)
		},
	)

	fmt.Println("backend3 running on :9003")

	http.ListenAndServe(
		":9003",
		nil,
	)
}