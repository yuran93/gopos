package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	err := godotenv.Load("/apps/gopos/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// dbUser := ""
	// dbPassword := ""
	// dbHost := ""
	// dbPort := ""
	// dbName := ""

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Unable to connect to the database")
	}

	return db
}
