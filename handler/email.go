package handler

import "github.com/gofiber/fiber/v2"

func EmailList(c *fiber.Ctx) error {
	var users []string

	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}
