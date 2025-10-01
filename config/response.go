package config

import "github.com/gofiber/fiber/v2"

// Success response helper
func Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  true,
		"message": "success",
		"data":    data,
	})
}

// Error response helper (dinamis)
func Error(c *fiber.Ctx, statusCode int, msg string, err interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"code":    statusCode,
		"status":  false,
		"message": msg,
		"error":   err,
	})
}
