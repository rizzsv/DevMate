package main

import (
	"devteamhub_be/internal/config"
	"devteamhub_be/internal/middleware"
	"devteamhub_be/internal/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	// "devteamhub_be/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	config.ConnectDatabase()

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	utils.InitLogger()
	utils.InfoLogger.Println("Starting DevMate Backend...")

	// routes.SetupPublicRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
