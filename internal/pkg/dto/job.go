package dto

import "time"

type JobUser struct {
	ID      uint64 `db:"id_user"`
	Name    string
	Surname string
	DOB     time.Time
}
