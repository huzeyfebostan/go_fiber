package routes

import (
	"github.com/gofiber/fiber"
	"github.com/huzeyfebostan/golang_practice/controllers"
	"github.com/huzeyfebostan/golang_practice/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUsers)
}
