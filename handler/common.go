package handler

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"success": false,
		"error":   "not found",
	})
}
