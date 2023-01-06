package handler

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/pgtype"
)

// GetEmails func gets all books.
// @Description Get all books.
// @Summary get all books
// @Tags Emails
// @Accept json
// @Produce json
// @Success 200 {array} model.Email
// @Router /v1/books [get]
func GetEmails(c *fiber.Ctx) error {
	return nil
}

func EmailList(c *fiber.Ctx) error {
	var users []string

	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}
