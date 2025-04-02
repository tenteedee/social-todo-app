package business

import (
	"context"
	"strings"

	"github.com/tenteedee/social-todo-app/modules/item/model"
)

type CreateItemStorage interface {
	Create(ctx context.Context, data model.TodoItemCreation) error
}

type createItemBusiness struct {
	storage CreateItemStorage
}

func NewCreateItemBusiness(storage CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{
		storage: storage,
	}
}

func (b *createItemBusiness) CreateItemBusiness(ctx context.Context, data model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrorTitleIsBlank
	}

	if err := b.storage.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
