package routes

import "github.com/gofiber/fiber/v2"

func SetupUserRoutes(router fiber.Router) {
	userRouter := router.Group("/user")
	userRouter.Get("/get-user", func(c *fiber.Ctx) error {
		return c.SendString("hello from user router")
	})
}
