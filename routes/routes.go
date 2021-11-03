package routes

import (
	"gopos/api"
	"os"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		dbName := os.Getenv("DB_NAME")
		return c.SendString("Selected database: " + dbName)
	})

	app.Post("/users", api.StoreUser)

	app.Get("/redis/get", api.GetRedisValue)
	app.Get("/redis/set", api.SetRedisValue)

	app.Get("/transactions", api.GetTransactions)

}
