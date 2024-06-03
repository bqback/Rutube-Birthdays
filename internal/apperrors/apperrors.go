package apperrors

import "errors"

var (
	ErrInvalidLoggingLevel = errors.New("invalid logging level")
)

var (
	ErrEnvNotFound       = errors.New("couldn't open .env file at provided path")
	ErrDatabasePWMissing = errors.New("database PW is missing from .env")
)

var (
	ErrCouldNotBuildQuery       = errors.New("couldn't build query to DB")
	ErrCouldNotBeginTransaction = errors.New("couldn't being a DB transaction")
	ErrCouldNotRollback         = errors.New("couldn't roll back the DB transaction")
	ErrEmptyResult              = errors.New("no entries returned on search")
)

var (
	ErrUserNotCreated  = errors.New("failed to create user")
	ErrUserNotSelected = errors.New("failed to select user")
	ErrWrongPassword   = errors.New("wrong password")
)
