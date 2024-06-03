package services

import (
	"birthdays/internal/services/monolith"
	"birthdays/internal/storages"
)

type Services struct {
	Auth IAuthService
	User IUserService
}

func NewServices(storages *storages.Storages) *Services {
	return &Services{
		Auth: monolith.NewAuthService(storages.Auth),
		User: monolith.NewUserService(storages.User),
	}
}
