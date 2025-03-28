package types

import "time"

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id"`
	Title       string     `json:"text" gorm:"column:title"`
	Images      string     `json:"images" gorm:"column:images"`
	Description string     `json:"description" gorm:"column:description"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (TodoItem) TableName() string {
	return "items"
}

type TodoItemCreation struct {
	Id          int    `json:"-" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}
