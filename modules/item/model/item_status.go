package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusNotYet ItemStatus = iota
	ItemStatusDoing
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatuses = [4]string{
	"not yet",
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
	if item == nil {
		return []byte("null"), nil
	}

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
