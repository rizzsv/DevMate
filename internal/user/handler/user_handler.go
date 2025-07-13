package handler

import (
	"devteamhub_be/internal/user/domain"
	"devteamhub_be/internal/user/usecase"
	"devteamhub_be/internal/user/validator"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service *usecase.UserService
}

func NewUserHandler(s *usecase.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var body domain.User
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validator.ValidateUserRegister(body); err != nil {
		return err
	}

	if err := h.Service.Register(&body); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"succes":  true,
		"message": "User registered successfully",
	})
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return err
	}
	return c.JSON(users)
}
