package main

import (
	"gopos/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// We'll skip the auto migration for now.
	// db := database.Connect()
	// db.AutoMigrate(&api.User{})

	app := fiber.New()
	routes.InitRoutes(app)
	app.Listen(":8022")

}
