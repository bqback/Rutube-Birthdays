package storages

import (
	"birthdays/internal/pkg/dto"
	"birthdays/internal/pkg/entities"
	"context"
)

type IAuthStorage interface {
	Auth(context.Context, dto.LoginInfo) (*dto.DBUser, error)
	Register(context.Context, dto.SignupInfo) (*entities.User, error)
}
