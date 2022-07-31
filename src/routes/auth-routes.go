package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/haris/go-rest/src/controllers"
)

func SetupAuthRoutes(router fiber.Router) {
	authRouter := router.Group("/auth")
	authRouter.Post("/create-admin", controllers.CreateAdmin)
}
