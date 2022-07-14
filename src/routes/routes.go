package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	SetupAuthRoutes(v1)
	SetupCategoriesRoutes(v1)
	SetupUserRoutes(v1)
	SetupProductRoutes(v1)
}
