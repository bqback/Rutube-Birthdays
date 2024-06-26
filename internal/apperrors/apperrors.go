package apperrors

import "errors"

var (
	ErrInvalidLoggingLevel = errors.New("invalid logging level")
)

var (
	ErrEnvNotFound       = errors.New("couldn't open .env file at provided path")
	ErrDatabasePWMissing = errors.New("database PW is missing from .env")
	ErrJWTSecretMissing  = errors.New("JWT secret is missing from .env")
)

var (
	ErrCouldNotParseClaims = errors.New("couldn't parse JWT claims")
	ErrTokenExpired        = errors.New("JWT expired")
	ErrInvalidIssuedTime   = errors.New("invalid IAT in the JWT")
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
	ErrUsernameTaken   = errors.New("username taken")
	ErrWrongPassword   = errors.New("wrong password")
)

var (
	ErrSubscriptionNotCreated    = errors.New("failed to subscribe to birthday")
	ErrSubscriptionNotDeleted    = errors.New("failed to unsubscribe from birthday")
	ErrSubscriptionAlreadyExists = errors.New("subscription already exists")
)

var (
	ErrEmailsNotSelected = errors.New("couldn't select emails of notifications recipients")
)

var (
	ErrNilContext      = errors.New("context is nil")
	ErrUserIDMissing   = errors.New("user ID is missing from context")
	ErrUsernameMissing = errors.New("username is missing from context")
	ErrParamMissing    = errors.New("one or more required parameters are missing from context")
	ErrLoggerMissing   = errors.New("logger missing from context")
)
