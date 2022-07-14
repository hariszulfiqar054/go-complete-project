package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	authRouter := router.Group("/auth")
	authRouter.Get("/create-user", func(c *fiber.Ctx) error {
		return c.SendString("hello from auth router")
	})
}
