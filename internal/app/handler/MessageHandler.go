package handler

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app/service"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app"
	"strconv"
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

func (m *MessageHandler) GetAllMsgFromChatID(c *fiber.Ctx) error {
	id := c.Params("id")
	chatId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid chat ID"})
	}

	messages, err := m.service.GetAllMessagesByChatID(uint(chatId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(messages)
}

func (m *MessageHandler) GetMessagesByUserID(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	msg, err := m.service.GetMessagesByUserID(uint(userId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(msg)
}
