package domain


type Taste struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
}