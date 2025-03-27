package types

import "time"

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"text"`
	Images      string     `json:"images"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
