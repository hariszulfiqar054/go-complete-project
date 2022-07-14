package routes

import "github.com/gofiber/fiber/v2"

func SetupCategoriesRoutes(router fiber.Router) {

	categoriesRouter := router.Group("/category")
	categoriesRouter.Get("/get-categories", func(c *fiber.Ctx) error {
		return c.SendString("Hello from categoires router")
	})

}
