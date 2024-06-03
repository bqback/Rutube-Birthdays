package handlers

import (
	"birthdays/internal/services"
)

type Handlers struct {
	AuthHandler
	UserHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		AuthHandler: *NewAuthHandler(services.Auth, services.User),
		UserHandler: *NewUserHandler(services.User),
	}
}

func NewAuthHandler(as services.IAuthService, us services.IUserService) *AuthHandler {
	return &AuthHandler{
		as: as,
		us: us,
	}
}

func NewUserHandler(us services.IUserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}
