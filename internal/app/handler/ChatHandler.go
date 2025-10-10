package handler

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app/service"
	"strconv"
)

type ChatHandler struct {
	service ChatService
}

func NewChatHandler(s ChatService) *ChatHandler {
	return &ChatHandler{service: s}
}

func (h *ChatHandler) CreateChat(c *fiber.Ctx) error {
	var req struct {
		Title string `json:"title"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	chat, err := h.service.CreateChat(req.Title)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(chat)

}

func (h *ChatHandler) GetChatByChatID(c *fiber.Ctx) error {
	id := c.Params("id")
	chatId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid chat ID"})
	}

	chat, err := h.service.GetChatByID(uint(chatId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Chat not found"})
	}
	return c.Status(fiber.StatusOK).JSON(chat)

}

func (h *ChatHandler) DeleteChat(c *fiber.Ctx) error {
	id := c.Params("id")
	chatId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid chat ID"})
	}

	err = h.service.DeleteChat(uint(chatId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Chat not found"})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)

}
