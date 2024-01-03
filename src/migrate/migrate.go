package main

import (
	"github.com/Daniel-Hermel-ML/crud-testing/cmd/configs"
	"github.com/Daniel-Hermel-ML/crud-testing/cmd/models"
)

func init() {
	configs.ConnectionDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Produto{})
}
