package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Notification represents the expected structure of incoming webhook data
type Notification struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
}

// webhookNotificationHandler processes POST requests to the webhook endpoint
func webhookNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Parse the JSON data
	var notification Notification
	if err := json.Unmarshal(body, &notification); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Log the notification data
	log.Printf("Webhook received: TransactionID=%s, Status=%s, Amount=%s, Currency=%s",
		notification.TransactionID, notification.Status, notification.Amount, notification.Currency)

	// Respond to the webhook
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Webhook received successfully")
}
