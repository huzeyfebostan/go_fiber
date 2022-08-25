package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/huzeyfebostan/golang_practice/database"
	"github.com/huzeyfebostan/golang_practice/models"
	"strconv"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
} //bütün rolleri getirir

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(&role)

}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
} //id'si girilen rolü getirir

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	var result interface{}
	database.DB.Table("role_permissions").Where("role_id", id).Delete(result)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
} //rolun bilgilerini güncelle

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(role)

	return nil
} //rol siler
