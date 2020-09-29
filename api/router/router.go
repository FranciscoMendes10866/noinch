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
	auth.Post("/create", handlers.SignUp)
	auth.Post("/get", handlers.SignIn)
	// manager routes
	manager := api.Group("/manager")
	manager.Get("/", guards.Protected(), handlers.GetRecords)
	manager.Post("/", guards.Protected(), handlers.CreateRecord)
	manager.Put("/:id", guards.Protected(), handlers.UpdateRecord)
	manager.Delete("/:id", guards.Protected(), handlers.DeleteRecord)
}
