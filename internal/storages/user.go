package storages

import (
	"birthdays/internal/pkg/entities"
	"context"
)

type IUserStorage interface {
	GetAll(context.Context) ([]*entities.User, error)
	Subscribe(context.Context, uint64, uint64) error
	Unsubscribe(context.Context, uint64, uint64) error
}
