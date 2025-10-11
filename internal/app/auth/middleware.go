package auth

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JWTProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("JWT_SECRET"),
		ErrorHandler: jwtError,
		ContextKey:   "user", // เก็บข้อมูล token ใน c.Locals("user")
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "unauthorized or invalid token",
	})
}