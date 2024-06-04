package monolith

import (
	"birthdays/internal/pkg/entities"
	"birthdays/internal/storages"
	"context"
)

type UserService struct {
	storage storages.IUserStorage
}

func NewUserService(storage storages.IUserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]*entities.User, error) {
	return s.storage.GetAll(ctx)
}

func (s *UserService) Subscribe(ctx context.Context, source uint64, id uint64) error {
	return s.storage.Subscribe(ctx, source, id)
}

func (s *UserService) Unsubscribe(ctx context.Context, source uint64, id uint64) error {
	return s.storage.Unsubscribe(ctx, source, id)
}
