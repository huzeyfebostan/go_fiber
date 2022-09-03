package database

import (
	"github.com/huzeyfebostan/golang_practice/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:hoze2626@/golang_practice"), &gorm.Config{})

	if err != nil {
		panic("Veritabanına bağlanılamadı !!!")
	}

	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
