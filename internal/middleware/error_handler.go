package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// fiber siap menangani error yang terjadi
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	log.Printf("Error %d: %s", code, message)

	return c.Status(code).JSON(fiber.Map{
		"succes": false,
		"error": message,
	})
}