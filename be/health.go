package main

import (
	"encoding/json"
	"net/http"
)

// HealthCheckResponse is the structure for health check response
type HealthCheckResponse struct {
	Status string `json:"status"`
}

// healthCheckHandler responds with the health status
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthCheckResponse{Status: "healthy"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
