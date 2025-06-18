package models

import "time"

type Todo struct {
	ID         int       `json:"id"`
	Task       string    `json:"task"`
	IsComplete bool      `json:"is_complete"`
	CreateAt   time.Time `json:"create_at"`
}
