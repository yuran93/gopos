package main

import (
	"gopos/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// We'll skip the auto migration for now.
	// db := database.Connect()
	// db.AutoMigrate(&api.User{})

	app := fiber.New()
	routes.InitRoutes(app)
	app.Listen(":8022")

}
