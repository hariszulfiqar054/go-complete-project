package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haris/go-rest/src/routes"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
