package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go_blog/database"
	"go_blog/router"
	"log"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Environment variables loaded successfully")
	database.ConnectionDB()
}

func main() {

	postgresDb, err := database.DBConnection.DB()
	if err != nil {
		panic("Failed to connect to postgres database!")
	}
	defer postgresDb.Close()

	app := fiber.New()

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":4000")
}
