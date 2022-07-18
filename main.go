package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haris/go-rest/src/config"
	"github.com/haris/go-rest/src/middleware"
	"github.com/haris/go-rest/src/routes"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	middleware.SetupMiddleware(app)
	app.Listen(":" + config.SERVER_PORT)
}
