package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/huzeyfebostan/golang_practice/database"
	"github.com/huzeyfebostan/golang_practice/routes"
)

func main() {
	database.Connect() //veritabanını bağlıyoruz

	app := fiber.New() //serveri başlatır

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app) //sayfaları ayarlar

	app.Listen(":8000") //8000 portunu dinler

}
