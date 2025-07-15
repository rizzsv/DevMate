package main

import (
	"devteamhub_be/config"
	"devteamhub_be/internal/user/domain"
	"devteamhub_be/internal/user/routes"
	"devteamhub_be/middleware"
	"devteamhub_be/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Connect ke database dan migrasi User
	config.ConnectDatabase()
	db := config.DB
	db.AutoMigrate(&domain.User{})

	// Init logger
	utils.InitLogger()
	utils.InfoLogger.Println("Starting DevMate Backend...")

	// Init Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Register route
	api := app.Group("/api")
	routes.UserRoutes(api.Group("/user")) // => /api/user/register, /api/user/login, etc

	// Start server
	log.Fatal(app.Listen(":3000"))
}
