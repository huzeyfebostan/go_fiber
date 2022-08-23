package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/huzeyfebostan/golang_practice/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "Ali",
	}

	user.LastName = "Veli"
	return c.JSON(user)
}
