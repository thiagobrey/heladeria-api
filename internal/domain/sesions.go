package domain

import "time"

type sesions struct {
	Id         int        `json:"id"`
	Token      string     `json:"token"`
	Timestamps Timestamps `json:"timestamps"`
}

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
