package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noinch-api/guards"
	"github.com/noinch-api/handlers"
)

// SetupRouter - setup router api
func SetupRouter(app *fiber.App) {
	// defining router prefix
	api := app.Group("/api/v1")
	// auth routes
	auth := api.Group("/auth")
	auth.Post("/create", handlers.CreateUser)
	auth.Post("/get", handlers.GetUser)
	// manager routes
	manager := api.Group("/manager")
	manager.Post("/create", guards.Protected(), handlers.CreateRecord)
}
