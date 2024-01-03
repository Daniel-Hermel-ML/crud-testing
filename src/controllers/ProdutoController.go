package controllers

import (
	"github.com/Daniel-Hermel-ML/crud-testing/cmd/configs"
	"github.com/Daniel-Hermel-ML/crud-testing/cmd/models"

	"github.com/gin-gonic/gin"
)

type ProdutoRequestBody struct {
	Titulo    string  `json:"titulo"`
	Preco     float32 `json:"preco"`
	Descricao string  `json:"descricao"`
	Imagem    string  `json:"imagem"`
}

// Start Criar Produto
func ProdutoCreate(c *gin.Context) {
	body := ProdutoRequestBody{}

	c.BindJSON(&body)

	produto := &models.Produto{Titulo: body.Titulo, Price: body.Preco, Descricao: body.Descricao,
		Image: body.Imagem}

	result := configs.DB.Create(&produto)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Falha ao inserir Produto"})
		return
	}
}

// End Criar Produto
// Start Buscar Produto
func ProdutoGet(c *gin.Context) {

	var produtos []models.Produto
	configs.DB.Find(&produtos)
	c.JSON(200, &produtos)
}

// End Buscar Produto
// Start Buscar Produto By ID
func ProdutoGetById(c *gin.Context) {

	id := c.Param("id")
	var produto models.Produto
	configs.DB.First(&produto, id)
	c.JSON(200, &produto)
}

// End Buscar Produto By ID
// Start Atualizar Produto
func ProdutoUpdate(c *gin.Context) {

	id := c.Param("id")
	var produto models.Produto
	configs.DB.First(&produto, id)

	body := ProdutoRequestBody{}
	c.BindJSON(&body)
	data := &models.Produto{Titulo: body.Titulo, Price: body.Preco, Descricao: body.Descricao,
		Image: body.Imagem}

	result := configs.DB.Model(&produto).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Falhou ao atualizar"})
		return
	}

	c.JSON(200, &produto)
}

// End Atualizar Produto
// Start Delete Produto
func ProdutoDelete(c *gin.Context) {
	id := c.Param("id")
	var produto models.Produto
	configs.DB.Delete(&produto, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}

//End Delete Produto
