package utils

import (
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/jubelio/go-logging/logging"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			logging.Errorf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(":" + os.Getenv("NODE_PORT")); err != nil {
		logging.Errorf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(":" + os.Getenv("NODE_PORT")); err != nil {
		logging.Errorf("Oops... Server is not running! Reason: %v", err)
	}
}