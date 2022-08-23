package routes

import (
	"github.com/gofiber/fiber"
	"github.com/huzeyfebostan/golang_practice/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
