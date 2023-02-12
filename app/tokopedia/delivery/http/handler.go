package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jubelio/chat-webhook/domain"
	"github.com/jubelio/go-logging/logging"
)

type handlerWebhook struct {
}

type HandlerTokopedia interface {
	ReceiveMessage(c *fiber.Ctx)
}

func NewTokopediaHandler() handlerWebhook {
	return handlerWebhook{}
}

func (h *handlerWebhook) ReceiveMessage(c *fiber.Ctx) error {
	var payload domain.TokopediaWebhookMsg
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.HTTPResponse{Message: err.Error(), Status: fiber.StatusBadRequest})
	}

	return c.Status(fiber.StatusOK).JSON(domain.HTTPResponse{Message: "success", Data: payload, Status: fiber.StatusOK})
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logging.Error(err.Error())
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
