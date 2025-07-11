package utils

import "github.com/gofiber/fiber/v2"

func NewHttpError(code int, message string) *fiber.Error {
	return fiber.NewError(code, message)
}