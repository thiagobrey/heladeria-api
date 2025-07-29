package domain

import (
	"gorm.io/gorm"
)

type Pedidos struct {
	gorm.Model
	UserId int     `json:"user_id"`
	Price  float32 `json:"price"`
	Code   int     `json:"code"`
	Tastes string  `json:"tastes"`
}

type PedidosRequest struct {
	UserId    int         `json:"user_id" binding:"required"`
	Productos []Productos `json:"productos" binding:"dive,required,count_validator"`
}

type Productos struct {
	Price  float32 `json:"price"`
	Code   int     `json:"code" binding:"required"`
	Tastes string  `json:"tastes" binding:"required"`
}

type PedidosResponse struct {
	User  User        `json:"user"`
	Items []Productos `json:"items"`
	Total float32     `json:"total"`
}



// {
// 	user:{

// 	},
// 	dia ( created_at):{
// 		Items: []Productos{
// 		}
// 	},
// 	dia ( created_at):{
// 		Items: []Productos{
// 		}
// 	},
// 	dia ( created_at):{
// 		Items: []Productos{
// 		}
// 	}
// }