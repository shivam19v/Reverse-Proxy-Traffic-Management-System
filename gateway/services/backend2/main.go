package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("backend2 healthy"))
		},
	)

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("backend2 FAILURE")

			http.Error(
				w,
				"backend2 failed",
				http.StatusInternalServerError,
			)
		},
	)

	fmt.Println("backend2 running on :9002")

	http.ListenAndServe(
		":9002",
		nil,
	)
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"net/http"
// 	"time"
// )

// func main() {

// 	rand.Seed(time.Now().UnixNano())

// 	http.HandleFunc(
// 		"/health",
// 		func(w http.ResponseWriter, r *http.Request) {

// 			// Health endpoint always healthy
// 			w.WriteHeader(http.StatusOK)
// 			w.Write([]byte("backend2 healthy"))
// 		},
// 	)

// 	http.HandleFunc(
// 		"/",
// 		func(w http.ResponseWriter, r *http.Request) {

// 			// Random failures
// 			random := rand.Intn(100)

// 			// 70% failure rate
// 			if random < 70 {

// 				fmt.Println("backend2 FAILURE")

// 				http.Error(
// 					w,
// 					"backend2 failed",
// 					http.StatusInternalServerError,
// 				)

// 				return
// 			}

// 			fmt.Println("backend2 SUCCESS")

// 			w.WriteHeader(http.StatusOK)

// 			w.Write([]byte(
// 				"response from backend2\n",
// 			))
// 		},
// 	)

// 	fmt.Println("backend2 running on :9002")

// 	http.ListenAndServe(
// 		":9002",
// 		nil,
// 	)
// }