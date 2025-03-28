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

type TodoItemCreation struct {
	Id          int    `json:"-" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

func (TodoItemCreation) TableName() string {
	return "items"
}
