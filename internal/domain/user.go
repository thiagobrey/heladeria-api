package domain

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Direction string    `json:"direction"`
	Pedidos   []Pedidos `json:"pedidos,omitempty" gorm:"foreignKey:UserId"`
}

type UserRequest struct {
	Name      string `json:"name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Direction string `json:"direction" binding:"required"`
}

type UserPedidosResponse struct {
	CreatedAt   string  `json:"created_at"`
	Description string  `json:"description"`
	Tastes      string  `json:"tastes"`
	Price       float64 `json:"price"`
}

type UserResponse struct {
	User    User          `json:"user"`
	Pedidos []PedidosUser `json:"pedidos,omitempty"`
}

type PedidosUser struct {
	Fecha string  `json:"fecha"`
	Items []Items `json:"items"`
}

type Items struct {
	Price  float32 `json:"price"`
	Code   int     `json:"code" binding:"required"`
	Tastes string  `json:"tastes" binding:"required"`
}
