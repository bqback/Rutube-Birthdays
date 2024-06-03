package entities

import "time"

type JWT struct {
	User  string
	Token string
}

type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
	DOB     time.Time
}
