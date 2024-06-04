package storages

import (
	"birthdays/internal/pkg/entities"
	"context"
)

type IUserStorage interface {
	GetAll(context.Context) ([]*entities.User, error)
}
