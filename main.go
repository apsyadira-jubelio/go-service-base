package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jubelio/chat-webhook/config"
	"github.com/jubelio/chat-webhook/utils"
)

func main() {
	// Load configuration
	godotenv.Load()

	app := fiber.New()
	app.Use(config.MiddlewareFiber)

	// Start the server
	utils.StartServerWithGracefulShutdown(app)
}
