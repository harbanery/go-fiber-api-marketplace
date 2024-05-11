package main

import (
	"gofiber-marketplace/src/configs"
	"gofiber-marketplace/src/helpers"
	"gofiber-marketplace/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	configs.InitDB()
	helpers.Migration()
	routes.Router(app)

	app.Listen(":3000")
}
