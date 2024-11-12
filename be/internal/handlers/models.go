package handler

import "encoding/xml"

// PaymentCreateReq is the structure for create payment request
type PaymentCreateReq struct {
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	TransactionId string `json:"transaction_id"`
}

// PaymentCreateResp is the structure for create payment response
type PaymentCreateResp struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	TransactionId string `json:"transaction_id"`
	UniqueID      string `json:"unique_id"`
	Timestamp     string `json:"timestamp"`
	RedirectURL   string `json:"redirect_url"`
	Message       string `json:"message,omitempty"`
	Success       bool   `json:"success"`
}

type EmerchantPayWpfPaymentReq struct {
	XMLName          xml.Name         `xml:"wpf_payment"`
	TransactionId    string           `xml:"transaction_id"`
	NotificationURL  string           `xml:"notification_url"`
	ReturnSuccessURL string           `xml:"return_success_url"`
	ReturnFailureURL string           `xml:"return_failure_url"`
	ReturnCancelURL  string           `xml:"return_cancel_url"`
	ReturnPendingURL string           `xml:"return_pending_url"`
	Amount           string           `xml:"amount"`
	Currency         string           `xml:"currency"`
	TransactionTypes TransactionTypes `xml:"transaction_types"`
	TerminalId       string           `xml:"terminal_id"`
}

type TransactionTypes struct {
	TransactionType TransactionType `xml:"transaction_type"`
}

type TransactionType struct {
	Name string `xml:"name,attr"`
}

type EmerchantPayWpfPaymentResp struct {
	XMLName       xml.Name `xml:"wpf_payment"`
	Status        string   `xml:"status"`
	UniqueID      string   `xml:"unique_id"`
	TransactionID string   `xml:"transaction_id"`
	Timestamp     string   `xml:"timestamp"`
	Amount        string   `xml:"amount"`
	Currency      string   `xml:"currency"`
	RedirectURL   string   `xml:"redirect_url"`
}
