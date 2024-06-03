package monolith

import (
	"birthdays/internal/pkg/dto"
	"birthdays/internal/pkg/entities"
	"birthdays/internal/storages"
	"context"
)

type AuthService struct {
	storage storages.IAuthStorage
}

func NewAuthService(storage storages.IAuthStorage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) Auth(ctx context.Context, info dto.LoginInfo) (*dto.DBUser, error) {
	return nil, nil
}

func (s *AuthService) Register(ctx context.Context, info dto.SignupInfo) (*entities.User, error) {
	return nil, nil
}
