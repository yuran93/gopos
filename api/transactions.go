package api

import (
	"gopos/database"

	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	ID         uint    `gorm:"primarykey"`
	Code       string  `json:"code"`
	SlipId     int     `json:"slip_id"`
	EmployeeId int     `json:"employee_id"`
	StatusId   int     `json:"status_id"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
}

func GetTransactions(c *fiber.Ctx) error {

	db := database.Connect()

	transactions := []Transaction{}

	db.Model(&Transaction{}).Find(&transactions)

	return c.JSON(transactions)

}
