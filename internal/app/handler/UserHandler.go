package handler

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app"
	. "github.com/saigrogyh/Real-Time-Chat-API/internal/app/service"
	"github.com/saigrogyh/Real-Time-Chat-API/internal/app/auth"
	// "strconv"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.service.Login(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	token, err := auth.GenerateToken(user.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

