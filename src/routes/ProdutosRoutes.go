package routes

import (
	"github.com/Daniel-Hermel-ML/crud-testing/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func ProdutoRouter(router *gin.Engine) {

	routes := router.Group("api/produtos")
	routes.POST("", controllers.ProdutoCreate)
	routes.GET("", controllers.ProdutoGet)
	routes.GET("/:id", controllers.ProdutoGetById)
	routes.PUT("/:id", controllers.ProdutoUpdate)
	routes.DELETE("/:id", controllers.ProdutoDelete)

}
