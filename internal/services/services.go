package services

import (
	"birthdays/internal/auth"
	"birthdays/internal/services/monolith"
	"birthdays/internal/storages"
)

type Services struct {
	Auth IAuthService
	User IUserService
}

func NewServices(storages *storages.Storages, manager *auth.AuthManager) *Services {
	return &Services{
		Auth: monolith.NewAuthJWTService(storages.Auth, manager),
		User: monolith.NewUserService(storages.User),
	}
}
