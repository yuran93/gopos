package api

import (
	"gopos/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func StoreUser(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	db := database.Connect()

	db.Create(&user)

	return c.SendStatus(200)
}

func GetUsers(c *fiber.Ctx) error {

	// db := database.Connect()

	// db.Query("SELECT * FROM users WHER")

	something := make(map[string]string)

	something["yuran"] = "fernando"
	something["test"] = c.Query("foo")

	return c.JSON(something)
}
