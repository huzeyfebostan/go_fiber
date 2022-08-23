package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:hoze2626@/golang_practice"), &gorm.Config{})

	if err != nil {
		panic("Veritabanına bağlanılamadı !!!")
	}
}
