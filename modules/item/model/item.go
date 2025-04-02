package model

import (
	"errors"

	"github.com/tenteedee/social-todo-app/common"
	"gorm.io/gorm"
)

var (
	ErrorTitleIsBlank  = errors.New("title cannot be blank")
	ErrorItemIsDeleted = errors.New("item is deleted")
)

type TodoItem struct {
	common.SQLModel
	Title       string      `json:"title" gorm:"column:title"`
	Images      string      `json:"images" gorm:"column:images"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string {
	return "items"
}

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

func (item *TodoItemCreation) BeforeCreate(tx *gorm.DB) error {
	if item.Status == nil {
		defaultStatus := ItemStatusNotYet
		item.Status = &defaultStatus
	}
	return nil
}

type TodoItemUpdate struct {
	Title       *string     `json:"title" gorm:"column:title"`
	Description *string     `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
