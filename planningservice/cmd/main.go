package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"planningservice/internal/api"
	"planningservice/internal/config"
)

func main() {

	app := fiber.New(fiber.Config{
		ServerHeader: "planning_service_v1",
		AppName:      "planning_service_v1",
	})

	// Register routes
	api.RegisterRoutes(app)

	// Listen on port
	err := app.Listen(config.EnvConfigs.LocalServerPort)
	if err != nil {
		log.Default().Println("Error while serving the api: ", err)
	}

}
