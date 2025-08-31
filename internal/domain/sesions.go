package domain

import "time"

type Sesions struct {
	Id         int        `json:"id"`
	Token      string     `json:"token"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Data       string     `json:"data"`
}

