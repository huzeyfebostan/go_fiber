package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/huzeyfebostan/golang_practice/database"
	"github.com/huzeyfebostan/golang_practice/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permission []models.Permission

	database.DB.Find(&permission)

	return c.JSON(permission)
}
