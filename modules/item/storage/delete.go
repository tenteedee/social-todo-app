package storage

import (
	"context"

	"github.com/tenteedee/social-todo-app/modules/item/model"
)

func (s *sql_store) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	if err := s.db.
		Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{"status": "deleted"}).Error; err != nil {
		return err
	}

	return nil
}
