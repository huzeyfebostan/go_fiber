package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/huzeyfebostan/golang_practice/database"
	"github.com/huzeyfebostan/golang_practice/models"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Preload("Role").Find(&users)

	return c.JSON(users)
} //bütün kullanıcıları getirir

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(&user)

} //yeni kullanıcı

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role ").Find(&user)

	return c.JSON(user)
} //id'si girilen kullanıcıyı getirir

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
} //kullanıcı bilgilerini günceller

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(user)

	return nil
} //kullanıcı siler
