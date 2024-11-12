package api

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/prime-trader/payment-processor/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Health check route
	app.Get("/health", handler.HealthCheck)

	// Payment routes
	app.Post("/payment/create", handler.PaymentCreate)

	// Webhook
	app.Post("/webhook-notification", handler.WebhookNotificationHandler)

}
