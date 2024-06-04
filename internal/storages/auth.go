package storages

import (
	"birthdays/internal/pkg/dto"
	"birthdays/internal/pkg/entities"
	"context"
)

type IAuthStorage interface {
	Create(context.Context, dto.SignupInfo) (*entities.User, error)
	GetByUsername(context.Context, string) (*dto.DBUser, error)
}
