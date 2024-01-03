package main

import (
	"github.com/Daniel-Hermel-ML/crud-testing/src/configs"
	"github.com/Daniel-Hermel-ML/crud-testing/src/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectionDB()
}

func main() {
	r := gin.Default()

	routes.ProdutoRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello CRUD - MELI",
		})
	})
	r.Run()

}
