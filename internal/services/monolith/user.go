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
