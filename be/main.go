package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go backend!")
	})

	// Register endpoints with logging middleware
	http.HandleFunc("/health", logRequest(healthCheckHandler))
	http.HandleFunc("/create-payment", logRequest(createPaymentHandler))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

// logRequest is a middleware function that logs incoming requests
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: Method=%s, URL=%s, Headers=%v", r.Method, r.URL, r.Header)
		next(w, r)
	}
}
