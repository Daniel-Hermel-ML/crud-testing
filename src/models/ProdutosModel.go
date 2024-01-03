package models

import (
	"time"

	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Id        uint
	Titulo    string
	Price     float32
	Descricao string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Produto) TableName() string {
	return "produtos"
}
