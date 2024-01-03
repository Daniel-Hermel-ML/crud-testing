package configs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	var err error
	dsn := "root:mysql1234@tcp(127.0.0.1:3306)/produtos_go?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados!")
	}
}
