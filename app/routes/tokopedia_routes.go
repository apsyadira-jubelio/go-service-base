package routes

import (
	"github.com/gofiber/fiber/v2"
	_tokopediaHttpDelivery "github.com/jubelio/chat-webhook/app/tokopedia/delivery/http"
)

func TokopediaWebhookRoutes(app *fiber.App) {
	handler := _tokopediaHttpDelivery.NewTokopediaHandler()

	route := app.Use("/tokopedia")
	route.Post("/notify/message", handler.ReceiveMessage)
}
