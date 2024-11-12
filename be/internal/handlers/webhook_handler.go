package handler

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/xml"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/prime-trader/payment-processor/config"
)

type Notification struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
}

type NotificationEcho struct {
	XMLName  xml.Name `xml:"notification_echo"`
	UniqueID string   `xml:"unique_id"`
}

// WebhookNotificationHandler processes POST requests to the webhook endpoint using Fiber
func WebhookNotificationHandler(c *fiber.Ctx) error {
	// Check if the method is POST
	if c.Method() != fiber.MethodPost {
		log.Printf("Error: Invalid request method: %s\n", c.Method())
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{"error": "Invalid request method"})
	}

	// Parse form data from request
	// if err := c.Request().ParseForm(); err != nil {
	// 	log.Printf("Error parsing form: %v\n", err)
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	// }

	// Extract the necessary fields from the webhook
	wpfTransactionID := c.FormValue("wpf_transaction_id")
	wpfStatus := c.FormValue("wpf_status")
	notificationType := c.FormValue("notification_type")
	paymentTransactionID := c.FormValue("payment_transaction_unique_id")
	signature := c.FormValue("signature")

	// Validate that the required fields are present
	if paymentTransactionID == "" || signature == "" {
		log.Printf("Error: Missing required parameters (paymentTransactionID: %s, signature: %s)\n", paymentTransactionID, signature)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required parameters"})
	}

	// Verify the signature (SHA-512 Hash of payment_transaction_unique_id + apiPassword)
	if !verifySignature(paymentTransactionID, signature) {
		log.Printf("Error: Invalid signature for transaction ID %s\n", paymentTransactionID)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid signature"})
	}

	// Log the extracted values (for debugging or processing)
	log.Printf("WPF Transaction ID: %s\n", wpfTransactionID)
	log.Printf("WPF Status: %s\n", wpfStatus)
	log.Printf("Notification Type: %s\n", notificationType)
	log.Printf("Payment Transaction ID: %s\n", paymentTransactionID)
	log.Printf("Signature: %s\n", signature)

	// Respond with the unique_id in XML to acknowledge the notification
	response := NotificationEcho{
		UniqueID: wpfTransactionID,
	}

	// Convert NotificationEcho struct to XML
	xmlResponse, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("Error encoding XML response: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create XML response"})
	}

	// Send the response with XML content type
	c.Set("Content-Type", "application/xml")
	return c.Status(fiber.StatusOK).Send(xmlResponse)
}

// verifySignature verifies the webhook signature using SHA-512 of the payment_transaction_unique_id + API password
func verifySignature(paymentTransactionID, receivedSignature string) bool {
	// Concatenate payment_transaction_unique_id and API password
	cfg := config.GetConfig()
	signatureString := paymentTransactionID + cfg.EmerchantPayPassword

	// Generate SHA-512 hash of the concatenated string
	hash := sha512.New()
	hash.Write([]byte(signatureString))
	expectedSignature := hex.EncodeToString(hash.Sum(nil))

	// Compare the generated signature with the received signature
	return strings.EqualFold(expectedSignature, receivedSignature)
}
