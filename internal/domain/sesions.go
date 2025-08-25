package domain

import "time"

type Sesions struct {
	Id         int        `json:"id"`
	Token      string     `json:"token"`
	Timestamps Timestamps `json:"timestamps"`
	Data       string     `json:"data"`
}

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
