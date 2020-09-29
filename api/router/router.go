package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noinch-api/handlers"
)

// SetupRouter - setup router api
func SetupRouter(app *fiber.App) {
	// defining router prefix
	api := app.Group("/api/v1")
	// auth routes
	auth := api.Group("/auth")
	auth.Post("/create", handlers.Create)
	auth.Post("/get", handlers.Get)
}
