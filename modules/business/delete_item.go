package business

import (
	"context"

	"github.com/tenteedee/social-todo-app/modules/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemBusiness struct {
	storage DeleteItemStorage
}

func NewDeleteItemBusiness(storage DeleteItemStorage) *deleteItemBusiness {
	return &deleteItemBusiness{
		storage: storage,
	}
}

func (b *deleteItemBusiness) NewDeleteItemBusiness(ctx context.Context, id int) error {

	item, err := b.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if item.Status != nil && *item.Status == model.ItemStatusDeleted {
		return model.ErrorItemIsDeleted
	}

	if err := b.storage.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
