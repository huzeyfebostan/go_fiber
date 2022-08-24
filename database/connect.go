package database

import (
	"github.com/huzeyfebostan/golang_practice/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:hoze2626@/golang_practice"), &gorm.Config{}) //veritabanını açar 1.parametre kullanıcı adı 2.parametre şifre 3.parametre veritabanı ismi

	if err != nil {
		panic("Veritabanına bağlanılamadı !!!")
	}

	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}) //veritabanında tablo oluşturuyor
}
