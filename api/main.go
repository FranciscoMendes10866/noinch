package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noinch-api/router"
)

func main() {
	app := fiber.New()
	router.SetupRouter(app)
	app.Listen(":4444")
}
