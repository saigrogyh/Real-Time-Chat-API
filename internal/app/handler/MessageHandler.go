package handler

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app/service"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app"
)

type MessageHandler struct {
	service MessageService
}

func NewMessageHandler(s MessageService) *MessageHandler {
	return &MessageHandler{service: s}
}

func (m *MessageHandler) SendMessage(c *fiber.Ctx) error {
	var req SendMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	msg, err := m.service.SendMessage(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(msg)

}