package domain

import (
	"gorm.io/gorm"
)

type Pedidos struct {
	gorm.Model
	UserId   int `json:"user_id" binding:"required"`
	Price    float32
	Cantidad int
	Tastes   string
}

type PedidosRequest struct {
	UserId    int         `json:"user_id" binding:"required"`
	Productos []Productos `json:"productos" binding:"required"`
}

type Productos struct {
	Price    float32 `json:"price" binding:"required"`
	Cantidad int     `json:"cantidad" binding:"required"`
	Tastes   string  `json:"tastes" binding:"required"`
}
