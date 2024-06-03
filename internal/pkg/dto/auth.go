package dto

import "time"

type LoginInfo struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type DBUser struct {
	ID           uint64
	Username     string
	PasswordHash string
}

type SignupInfo struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Email    string    `json:"email" validate:"email"`
	DOB      time.Time `json:"dob" validate:"datetime"`
}
