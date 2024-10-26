package handlers

import (
	"app/cmd/internal/helpers"

	"github.com/gofiber/fiber/v2"
)

func RequestEmailHandler(c *fiber.Ctx) error {
	err := helpers.SendEmail("viniciuscontato234@gmail.com", "Assunto Teste", "Teste")

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString("Email sent successfully")
}
