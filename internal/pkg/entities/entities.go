package entities

import "time"

type JWT struct {
	User  string
	Token string
}

type User struct {
	ID      uint64 `db:"id_user"`
	Name    string
	Surname string
	Email   string
	DOB     time.Time
}

type Job struct {
	Date       time.Time
	SourceID   uint64
	SourceName string
}
