package domain

type Admin struct {
	Id        int    `json:"id"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Phone     string `json:"phone" binding:"required"`
	Birthdate string `json:"birthdate" binding:"required"`
}
