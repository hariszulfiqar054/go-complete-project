package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/haris/go-rest/src/dto"
	"github.com/haris/go-rest/src/utils"
)

type User struct {
	Name string
}

func CreateAdmin(c *fiber.Ctx) error {
	user := new(dto.CreateAdmin)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	fmt.Println("user", user)

	res := utils.New(true, "User created", user)
	return c.JSON(res)
	// services.CreateAdmin()
}
