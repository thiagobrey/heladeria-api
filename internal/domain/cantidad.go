package domain

type Cantidad struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	CountTastes int     `json:"count_tastes"`
	Code        int     `json:"code"`
	Price       float64 `json:"price"`
}

type CantidadRequest struct {
	Description string  `json:"description" binding:"required"`
	CountTastes int     `json:"count_tastes" binding:"required"`
	Code        int     `json:"code" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}