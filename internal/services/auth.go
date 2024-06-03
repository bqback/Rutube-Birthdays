package services

import (
	"birthdays/internal/pkg/dto"
	"birthdays/internal/pkg/entities"
	"context"
)

type IAuthService interface {
	Auth(context.Context, dto.LoginInfo) (*entities.JWT, error)
	Register(context.Context, dto.SignupInfo) (*dto.SignupResponse, error)
}
