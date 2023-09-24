package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// RegisterRoutes registers all routes for the API
func RegisterRoutes(app *fiber.App) {

	routes := app.Group("/api/v1/todo")

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	routes.Get("/planning", Planning) // Planning is a handler for planning

}
