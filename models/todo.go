package models

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	IsComplete bool     `json:"is_complete"`
	CreatedAt time.Time `json:"created_at"`
}
