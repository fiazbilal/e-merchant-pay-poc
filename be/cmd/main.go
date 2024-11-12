package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prime-trader/payment-processor/api"
	"github.com/prime-trader/payment-processor/config"
)

func main() {
	// Load configurations
	cfg := config.GetConfig()

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	api.SetupRoutes(app)

	// Log that the server is starting
	log.Printf("Server is starting at port: %s\n", cfg.Port)

	// Start the server
	log.Fatal(app.Listen(":" + cfg.Port))
}
