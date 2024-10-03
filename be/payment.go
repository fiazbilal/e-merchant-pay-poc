package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PaymentRequest is the structure for create payment request
type PaymentRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// CreatePaymentResponse is the structure for create payment response
type CreatePaymentResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// createPaymentHandler handles payment creation requests
func createPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var paymentRequest PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Here you would typically process the payment with the provided data
	// For demonstration, we will just simulate a successful payment response
	response := CreatePaymentResponse{
		Message: fmt.Sprintf("Payment of %.2f %s created successfully!", paymentRequest.Amount, paymentRequest.Currency),
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
