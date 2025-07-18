package domain

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Direction string `json:"direction"`
}

type UserRequest struct {
	Name      string `json:"name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Direction string `json:"direction" binding:"required"`
}
