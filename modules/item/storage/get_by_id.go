package storage

import (
	"context"

	"github.com/tenteedee/social-todo-app/modules/item/model"
)

func (s *sql_store) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data = model.TodoItem{}

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
