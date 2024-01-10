package config

import (
	"MondTemplate/app/middleware/database"
	"MondTemplate/app/middleware/utilities"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Initialize Golang Fiber
var App = fiber.New(fiber.Config{
	UnescapePath: true,
})

// DB Connectivity configuration
func CreateConnection() {
	username, dbConfigErr := utilities.Decrypt(utilities.GetEnv("POSTGRES_USERNAME"), utilities.GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	password, dbConfigErr := utilities.Decrypt(utilities.GetEnv("POSTGRES_PASSWORD"), utilities.GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	host, dbConfigErr := utilities.Decrypt(utilities.GetEnv("POSTGRES_HOST"), utilities.GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	dbName, dbConfigErr := utilities.Decrypt(utilities.GetEnv("DATABASE_NAME"), utilities.GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}

	fmt.Println("username: ", username)
	fmt.Println("password: ", password)
	fmt.Println("host: ", host)
	fmt.Println("dbName: ", dbName)
	// database.PostgreSQLConnect(
	// 	utilities.GetEnv("POSTGRES_USERNAME"),
	// 	utilities.GetEnv("POSTGRES_PASSWORD"),
	// 	utilities.GetEnv("POSTGRES_HOST"),
	// 	utilities.GetEnv("DATABASE_NAME"),
	// 	utilities.GetEnv("POSTGRES_PORT"),
	// 	utilities.GetEnv("POSTGRES_SSL_MODE"),
	// 	utilities.GetEnv("POSTGRES_TIMEZONE"),
	// )
	database.PostgreSQLConnect(
		username,
		password,
		host,
		dbName,
		utilities.GetEnv("POSTGRES_PORT"),
		utilities.GetEnv("POSTGRES_SSL_MODE"),
		utilities.GetEnv("POSTGRES_TIMEZONE"),
	)
}

func CreateTable(tableName string) error {
	err := database.DBConn.AutoMigrate(tableName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
