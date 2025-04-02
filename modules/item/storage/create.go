package storage

import (
	"context"

	"github.com/tenteedee/social-todo-app/modules/item/model"
)

func (s *sql_store) Create(ctx context.Context, data model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
