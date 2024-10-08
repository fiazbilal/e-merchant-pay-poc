package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	// Define a simple hello handler for the root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go backend!")
	})

	// Register endpoints with logging middleware
	http.HandleFunc("/health", logRequest(healthCheckHandler))
	http.HandleFunc("/create-payment", logRequest(createPaymentHandler))
	http.HandleFunc("/webhook-notification", logRequest(webhookNotificationHandler)) // New webhook notification endpoint

	// CORS middleware to allow requests from React frontend
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),  // Allow requests from your frontend's origin
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), // Allow GET, POST, OPTIONS methods
		handlers.AllowedHeaders([]string{"Content-Type"}),           // Allow Content-Type header
	)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", corsHandler(http.DefaultServeMux))
}

// logRequest is a middleware function that logs incoming requests
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: Method=%s, URL=%s", r.Method, r.URL)
		next(w, r)
	}
}

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

// PaymentRequest is the structure for create payment request
type PaymentRequest struct {
	TransactionID    string            `json:"transaction_id"`
	Amount           float64           `json:"amount"`
	Currency         string            `json:"currency"`
	NotificationURL  string            `json:"notification_url"`
	ReturnSuccessURL string            `json:"return_success_url"`
	ReturnFailureURL string            `json:"return_failure_url"`
	TransactionTypes []TransactionType `json:"transaction_types"`
}

// CreatePaymentResponse is the structure for create payment response
type CreatePaymentResponse struct {
	RedirectURL string `json:"redirect_url,omitempty"`
	Message     string `json:"message,omitempty"`
	Success     bool   `json:"success"`
}

type WpfPayment struct {
	XMLName       xml.Name `xml:"wpf_payment"`
	TransactionID string   `xml:"transaction_id"`
	// Usage            string   `xml:"usage"`
	// Description      string   `xml:"description"`
	NotificationURL  string `xml:"notification_url"`
	ReturnSuccessURL string `xml:"return_success_url"`
	ReturnFailureURL string `xml:"return_failure_url"`
	ReturnCancelURL  string `xml:"return_cancel_url"`
	ReturnPendingURL string `xml:"return_pending_url"`
	Amount           string `xml:"amount"`
	Currency         string `xml:"currency"`
	// ConsumerID       string   `xml:"consumer_id"`
	// CustomerEmail    string   `xml:"customer_email"`
	// CustomerPhone    string   `xml:"customer_phone"`
	// RememberCard     string   `xml:"remember_card"`
	// Lifetime         string   `xml:"lifetime"`
	TransactionTypes TransactionTypes `xml:"transaction_types"`
	TerminalID       string           `xml:"terminal_id"` // Add TerminalID field here

}

type TransactionTypes struct {
	TransactionType TransactionType `xml:"transaction_type"`
}

type TransactionType struct {
	Name string `xml:"name,attr"`
}

type WpfResponse struct {
	TransactionType  string `xml:"transaction_type"`
	Status           string `xml:"status"`
	UniqueID         string `xml:"unique_id"`
	TransactionID    string `xml:"transaction_id"`
	TechnicalMessage string `xml:"technical_message"`
	Message          string `xml:"message"`
	RedirectURL      string `xml:"redirect_url"`
	Amount           string `xml:"amount"`
	Currency         string `xml:"currency"`
	TerminalID       string `xml:"terminal_id"` // Add TerminalID field here
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

	// Print each field explicitly
	fmt.Println("Transaction ID:", paymentRequest.TransactionID)
	fmt.Println("Amount:", paymentRequest.Amount)
	fmt.Println("Currency:", paymentRequest.Currency)
	fmt.Println("Notification URL:", paymentRequest.NotificationURL)
	fmt.Println("Return Success URL:", paymentRequest.ReturnSuccessURL)
	fmt.Println("Return Failure URL:", paymentRequest.ReturnFailureURL)

	// Prepare the WpfPayment struct
	wpfPayment := WpfPayment{
		XMLName:       xml.Name{Local: "wpf_payment"},
		TransactionID: "31",
		// Usage:            paymentRequest.Usage,
		// Description:      paymentRequest.Description,
		NotificationURL:  paymentRequest.NotificationURL,
		ReturnSuccessURL: paymentRequest.ReturnSuccessURL,
		ReturnFailureURL: paymentRequest.ReturnFailureURL,
		ReturnCancelURL:  paymentRequest.ReturnSuccessURL,
		ReturnPendingURL: paymentRequest.ReturnSuccessURL,
		Amount:           "1000", // convert to string
		Currency:         "EUR",
		// ConsumerID:       paymentRequest.ConsumerID,
		// CustomerEmail:    paymentRequest.CustomerEmail,
		// CustomerPhone:    paymentRequest.CustomerPhone,
		// RememberCard:     strconv.FormatBool(paymentRequest.RememberCard),
		// Lifetime:         strconv.Itoa(paymentRequest.Lifetime),
		TransactionTypes: TransactionTypes{
			TransactionType: TransactionType{
				Name: "sale3d",
			},
		},
		TerminalID: "bba8038257ec6f93b7ae45dcd0bbd8a809e007bb",
	}

	// Convert WpfPayment struct to XML
	xmlPayload, err := xml.MarshalIndent(wpfPayment, "", "  ")
	if err != nil {
		http.Error(w, "Failed to create XML payload", http.StatusInternalServerError)
		return
	}

	// Construct the correct URL
	locale := "en" // Default locale
	// if paymentRequest.Locale != "" {
	// 	locale = paymentRequest.Locale
	// }
	apiURL := fmt.Sprintf("https://staging.wpf.emerchantpay.net/%s/wpf", locale)

	// Make the HTTP request to emerchantpay
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(xmlPayload))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Set your emerchantpay API credentials here
	req.SetBasicAuth("a9c05be020fc8a49fc736a925d48743edaffc268", "8ef19de3eb4ba5a9f7193c1762f27eafe7bb59d9")
	req.Header.Set("Content-Type", "application/xml")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	fmt.Println("Body: ", string(body))

	var wpfResponse WpfResponse
	err = xml.Unmarshal(body, &wpfResponse)
	if err != nil {
		http.Error(w, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	fmt.Println("wpfResponse: ", wpfResponse)

	// Check if the response contains a redirect_url
	if wpfResponse.RedirectURL == "" {
		http.Error(w, "Payment creation failed: No redirect URL provided", http.StatusInternalServerError)
		return
	}

	// Create the response object to send back to the client
	response := CreatePaymentResponse{
		RedirectURL: wpfResponse.RedirectURL,
		Message:     "Payment session created successfully. Redirect the customer to the provided URL.",
		Success:     true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

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

	fmt.Println("Webhook body: ", string(body))

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
