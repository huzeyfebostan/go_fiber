package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/huzeyfebostan/go_fiber/database"
	"github.com/huzeyfebostan/go_fiber/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permission []models.Permission

	database.DB.Find(&permission)

	return c.JSON(permission)
}
