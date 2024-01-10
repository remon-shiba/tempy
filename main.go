package main

import (
	"MondTemplate/app/config"
	"MondTemplate/app/middleware/utilities"
	"MondTemplate/app/routers"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Configure application CORS
	config.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Declare & initialize logger
	config.App.Use(logger.New())

	// Declare & initialize routes
	routers.SetupPublicRoutes(config.App)

	// Environment cofiguration
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}
	envi := os.Getenv("ENVIRONMENT")

	err = godotenv.Load(fmt.Sprintf("./environment/.env-%v", envi)) //
	if err != nil {
		log.Fatal("Error Loading Env File: ", err)
	}

	fmt.Println("ENV:", os.Getenv("PORT"))

	// Database connection
	config.CreateConnection()
	// Serve the application
	if utilities.GetEnv("SSL") == "enabled" {
		log.Fatal(config.App.ListenTLS(
			fmt.Sprintf(":%s", utilities.GetEnv("PORT")),
			utilities.GetEnv("SSL_CERTIFICATE"),
			utilities.GetEnv("SSL_KEY"),
		))
	} else {
		err := config.App.Listen(fmt.Sprintf(":%s", utilities.GetEnv("PORT")))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
