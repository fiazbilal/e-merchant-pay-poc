package handler

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/fiazbilal/e-merchant-pay-poc/config"
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	// Return a simple message or status
	return c.JSON(fiber.Map{
		"status":  "UP",
		"message": "Payment service is running",
	})
}

// PaymentCreate handles payment creation requests using Fiber
func PaymentCreate(c *fiber.Ctx) error {
	// Decode the incoming JSON request
	var body PaymentCreateReq
	if err := c.BodyParser(&body); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	// Get config
	cfg := config.GetConfig()

	// Prepare the WpfPayment struct
	eMerchantPayReqBody := EmerchantPayWpfPaymentReq{
		XMLName:          xml.Name{Local: "wpf_payment"},
		TransactionId:    body.TransactionId,
		NotificationURL:  cfg.NotificationURL,
		ReturnSuccessURL: cfg.ReturnSuccessURL,
		ReturnFailureURL: cfg.ReturnFailureURL,
		ReturnCancelURL:  cfg.ReturnCancelURL,
		ReturnPendingURL: cfg.ReturnPendingURL,
		Amount:           strconv.FormatInt(body.Amount, 10),
		Currency:         body.Currency,
		TransactionTypes: TransactionTypes{
			TransactionType: TransactionType{
				Name: "sale3d",
			},
		},
		TerminalId: cfg.TerminalId, // Assuming you have this in the environment
	}

	// Convert WpfPayment struct to XML
	xmlPayload, err := xml.MarshalIndent(eMerchantPayReqBody, "", "  ")
	if err != nil {
		log.Printf("Error creating XML payload: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create XML payload"})
	}

	// Make the HTTP request to emerchantpay
	client := &http.Client{}
	req, err := http.NewRequest("POST", cfg.EmerchantPayWpfURL, bytes.NewBuffer(xmlPayload))
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create request"})
	}

	// Set your emerchantpay API credentials here
	req.SetBasicAuth(cfg.EmerchantPayUsername, cfg.EmerchantPayPassword)
	req.Header.Set("Content-Type", "application/xml")

	eMerchantPayResp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send request"})
	}
	defer eMerchantPayResp.Body.Close()

	eMerchantPayRespRawBody, err := io.ReadAll(eMerchantPayResp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read response"})
	}

	// Parse the XML response
	var eMerchantPayRespBody EmerchantPayWpfPaymentResp
	err = xml.Unmarshal(eMerchantPayRespRawBody, &eMerchantPayRespBody)
	if err != nil {
		log.Printf("Error parsing XML response: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse response"})
	}

	// Check if the response is error
	if eMerchantPayRespBody.Status == "error" || eMerchantPayRespBody.RedirectURL == "" {
		log.Println("Payment creation failed error: ", string(eMerchantPayRespRawBody))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Payment creation failed"})
	}

	// Create the response object to send back to the client
	response := PaymentCreateResp{
		Amount:        eMerchantPayRespBody.Amount,
		Currency:      eMerchantPayRespBody.Currency,
		TransactionId: eMerchantPayRespBody.TransactionID,
		UniqueID:      eMerchantPayRespBody.UniqueID,
		Timestamp:     eMerchantPayRespBody.Timestamp,
		RedirectURL:   eMerchantPayRespBody.RedirectURL,
		Message:       "Payment session created successfully. Redirect the customer to the provided URL.",
		Success:       true,
	}

	// Send the response back to the client
	return c.Status(fiber.StatusCreated).JSON(response)
}
