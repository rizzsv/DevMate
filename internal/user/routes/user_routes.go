package routes

import (
	"devteamhub_be/config"
	"devteamhub_be/internal/user/handler"
	"devteamhub_be/internal/user/repository"
	"devteamhub_be/internal/user/usecase"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(routes fiber.Router) {
	repo := repository.NewUserRepository(config.DB)
	service := usecase.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	routes.Post("/register", handler.Register)
	routes.Get("/user", handler.GetAll)

	routes.Post("/login", handler.Login)
}
