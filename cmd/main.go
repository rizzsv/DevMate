package main

import (
	"devteamhub_be/config"
	"devteamhub_be/internal/user/domain"
	"devteamhub_be/internal/user/routes"
	"devteamhub_be/middleware"
	"devteamhub_be/utils"
	"log"
	_ "net/http/pprof"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	// "devteamhub_be/internal/routes"
)

func init() {
	go func() {
		log.Println(http.ListenAndServe("localhost:3000", nil))
	}()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	config.ConnectDatabase()
	db := config.DB 
	db.AutoMigrate(&domain.User{})

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	api := app.Group("/api")
	routes.UserRoutes(api.Group("/user"))

	utils.InitLogger()
	utils.InfoLogger.Println("Starting DevMate Backend...")
	

	// routes.SetupPublicRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
