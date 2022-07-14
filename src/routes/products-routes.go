package routes

import "github.com/gofiber/fiber/v2"

func SetupProductRoutes(router fiber.Router) {
	productRouter := router.Group("/product")
	productRouter.Get("/get-products", func(c *fiber.Ctx) error {
		return c.SendString("Hello from products")
	})
}
