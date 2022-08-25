package controllers

import (
	"github.com/gofiber/fiber"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	file := form.File["image"]
	filename := ""

	for _, file := range file {
		filename = file.Filename

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:8000/api/uploads/" + filename,
	})
}
