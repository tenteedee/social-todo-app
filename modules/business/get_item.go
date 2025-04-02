package business

import (
	"context"

	"github.com/tenteedee/social-todo-app/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type getItemBusiness struct {
	storage GetItemStorage
}

func NewGetItemBusiness(storage GetItemStorage) *getItemBusiness {
	return &getItemBusiness{
		storage: storage,
	}
}

func (b *getItemBusiness) NewGetItemBusiness(ctx context.Context, id int) (*model.TodoItem, error) {

	data, err := b.storage.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
