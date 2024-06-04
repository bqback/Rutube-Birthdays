package services

import (
	"birthdays/internal/pkg/entities"
	"context"
)

type IUserService interface {
	GetAll(context.Context) ([]*entities.User, error)
}
