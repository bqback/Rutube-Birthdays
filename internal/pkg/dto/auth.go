package dto

import (
	"birthdays/internal/pkg/entities"
	"time"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DBUser struct {
	ID           uint64
	Username     string
	PasswordHash string
}

type TokenInfo struct {
	ID       uint64
	Username string
}

type SignupInfo struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email" validate:"email"`
	DOB      time.Time `json:"dob" validate:"datetime"`
}

type SignupResponse struct {
	Token string
	entities.User
}
