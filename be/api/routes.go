package api

import (
	handler "github.com/fiazbilal/e-merchant-pay-poc/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health check route
	app.Get("/health", handler.HealthCheck)

	// Payment routes
	app.Post("/payment/create", handler.PaymentCreate)

	// Webhook
	app.Post("/webhook-notification", handler.WebhookNotificationHandler)

}
