package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id"`
	Title       string     `json:"title" gorm:"column:title"`
	Images      string     `json:"images" gorm:"column:images"`
	Description string     `json:"description" gorm:"column:description"`
	Status      ItemStatus `json:"status" gorm:"column:status"`
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

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}

type ItemStatus int

const (
	ItemStatusNotYet ItemStatus = iota
	ItemStatusDoing
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatuses = [4]string{
	"not_yet",
	"doing",
	"done",
	"deleted",
}

func (item *ItemStatus) String() string {
	return allItemStatuses[*item]
}

func ParseStringToItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatuses {
		if allItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), fmt.Errorf("invalid item status: %s", s)
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		// return errors.New(fmt.Sprintf("failed to scan item status: %s", value))
		return fmt.Errorf("failed to scan item status: %s", value)
	}

	v, err := ParseStringToItemStatus(string(bytes))
	if err != nil {
		return err
	}

	*item = v
	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemValue, err := ParseStringToItemStatus(str)
	if err != nil {
		return err
	}
	*item = itemValue
	return nil
}
