package monolith

import "birthdays/internal/storages"

type UserService struct {
	storage storages.IUserStorage
}

func NewUserService(storage storages.IUserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}
