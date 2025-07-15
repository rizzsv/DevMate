package routes

import (
	"devteamhub_be/config"
	"devteamhub_be/internal/user/handler"
	"devteamhub_be/internal/user/repository"
	"devteamhub_be/internal/user/usecase"
    "github.com/gofiber/jwt/v3"
	"os"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	repo := repository.NewUserRepository(config.DB)
	service := usecase.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)

	protected := router.Group("/", jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	protected.Get("/getUser", handler.GetAll)
}

