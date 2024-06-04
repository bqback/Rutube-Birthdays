package services

import (
	"birthdays/internal/auth"
	"birthdays/internal/services/monolith"
	"birthdays/internal/storages"

	"github.com/go-co-op/gocron/v2"
)

type Services struct {
	Auth IAuthService
	User IUserService
	Job  IJobService
}

func NewServices(storages *storages.Storages, manager *auth.AuthManager, scheduler gocron.Scheduler) *Services {
	return &Services{
		Auth: monolith.NewAuthJWTService(storages.Auth, manager),
		User: monolith.NewUserService(storages.User),
		Job:  monolith.NewJobService(storages.Job, scheduler),
	}
}
