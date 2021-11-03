package routes

import (
	"gopos/api"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Post("/users", api.StoreUser)

	app.Get("/redis/get", api.GetRedisValue)
	app.Get("/redis/set", api.SetRedisValue)

	app.Get("/transactions", api.GetTransactions)

}
