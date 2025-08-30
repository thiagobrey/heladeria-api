package domain

import "time"

type Sesions struct {
	Id         int        `json:"id"`
	Token      string     `json:"token"`
	Timestamps Timestamps `json:"timestamps" gorm:"embedded"`
	Data       string     `json:"data"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
