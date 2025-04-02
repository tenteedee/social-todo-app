package storage

import (
	"context"

	"github.com/tenteedee/social-todo-app/modules/item/model"
)

func (s *sql_store) UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error {

	if err := s.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
