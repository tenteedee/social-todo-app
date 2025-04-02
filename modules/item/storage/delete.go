package storage

import (
	"context"
)

func (s *sql_store) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	if err := s.db.Where(cond).Updates(map[string]interface{}{"status": "deleted"}).Error; err != nil {
		return err
	}

	return nil
}
