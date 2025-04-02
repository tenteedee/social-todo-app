package business

import (
	"context"

	"github.com/tenteedee/social-todo-app/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, data *model.TodoItemUpdate) error
}

type updateItemBusiness struct {
	storage UpdateItemStorage
}

func NewUpdateItemBusiness(storage UpdateItemStorage) *updateItemBusiness {
	return &updateItemBusiness{
		storage: storage,
	}
}

func (b *updateItemBusiness) NewUpdateItemBusiness(ctx context.Context, id int, data *model.TodoItemUpdate) error {

	item, err := b.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if item.Status != nil && *item.Status == model.ItemStatusDeleted {
		return model.ErrorItemIsDeleted
	}

	if err := b.storage.UpdateItem(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}
