package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PaymentRequest is the structure for create payment request
type PaymentRequest struct {
	TransactionID    string  `json:"transaction_id"`
	Usage            string  `json:"usage"`
	Description      string  `json:"description"`
	NotificationURL  string  `json:"notification_url"`
	ReturnSuccessURL string  `json:"return_success_url"`
	ReturnFailureURL string  `json:"return_failure_url"`
	ReturnCancelURL  string  `json:"return_cancel_url"`
	ReturnPendingURL string  `json:"return_pending_url"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
	ConsumerID       string  `json:"consumer_id"`
	CustomerEmail    string  `json:"customer_email"`
	CustomerPhone    string  `json:"customer_phone"`
	RememberCard     string  `json:"remember_card"`
	Lifetime         string  `json:"lifetime"`
}

// CreatePaymentResponse is the structure for create payment response
type CreatePaymentResponse struct {
	RedirectURL string `json:"redirect_url,omitempty"`
	Message     string `json:"message,omitempty"`
	Success     bool   `json:"success"`
}

// WpfPayment represents the structure of the payment request for emerchantpay
type WpfPayment struct {
	XMLName          xml.Name `xml:"wpf_payment"`
	TransactionID    string   `xml:"transaction_id"`
	Usage            string   `xml:"usage"`
	Description      string   `xml:"description"`
	NotificationURL  string   `xml:"notification_url"`
	ReturnSuccessURL string   `xml:"return_success_url"`
	ReturnFailureURL string   `xml:"return_failure_url"`
	ReturnCancelURL  string   `xml:"return_cancel_url"`
	ReturnPendingURL string   `xml:"return_pending_url"`
	Amount           string   `xml:"amount"`
	Currency         string   `xml:"currency"`
	ConsumerID       string   `xml:"consumer_id"`
	CustomerEmail    string   `xml:"customer_email"`
	CustomerPhone    string   `xml:"customer_phone"`
	RememberCard     string   `xml:"remember_card"`
	Lifetime         string   `xml:"lifetime"`
}

// WpfResponse represents the structure of the payment response from emerchantpay
type WpfResponse struct {
	XMLName     xml.Name `xml:"wpf_payment"`
	RedirectURL string   `xml:"redirect_url"`
	UniqueID    string   `xml:"unique_id"`
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

	// Prepare the WpfPayment struct
	wpfPayment := WpfPayment{
		TransactionID:    paymentRequest.TransactionID,
		Usage:            paymentRequest.Usage,
		Description:      paymentRequest.Description,
		NotificationURL:  paymentRequest.NotificationURL,
		ReturnSuccessURL: paymentRequest.ReturnSuccessURL,
		ReturnFailureURL: paymentRequest.ReturnFailureURL,
		ReturnCancelURL:  paymentRequest.ReturnCancelURL,
		ReturnPendingURL: paymentRequest.ReturnPendingURL,
		Amount:           fmt.Sprintf("%.2f", paymentRequest.Amount), // convert to string
		Currency:         paymentRequest.Currency,
		ConsumerID:       paymentRequest.ConsumerID,
		CustomerEmail:    paymentRequest.CustomerEmail,
		CustomerPhone:    paymentRequest.CustomerPhone,
		RememberCard:     paymentRequest.RememberCard,
		Lifetime:         paymentRequest.Lifetime,
	}

	// Convert WpfPayment struct to XML
	xmlPayload, err := xml.MarshalIndent(wpfPayment, "", "  ")
	if err != nil {
		http.Error(w, "Failed to create XML payload", http.StatusInternalServerError)
		return
	}

	// Make the HTTP request to emerchantpay
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://staging.gate.e-comprocessing.net/process/wpf", bytes.NewBuffer(xmlPayload))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Set your emerchantpay API credentials here
	req.SetBasicAuth("your_username", "your_password")
	req.Header.Set("Content-Type", "application/xml")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	var wpfResponse WpfResponse
	err = xml.Unmarshal(body, &wpfResponse)
	if err != nil {
		http.Error(w, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	// Create the response object to send back to the client
	response := CreatePaymentResponse{
		RedirectURL: wpfResponse.RedirectURL,
		Message:     "Payment created successfully!",
		Success:     true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
