package logging

import (
	"birthdays/internal/config"

	"github.com/go-chi/httplog/v2"
)

func Setup(config *config.LoggingConfig) (*httplog.Logger, error) {
	httpLogger := httplog.NewLogger("birthdays-api", httplog.Options{
		LogLevel:        config.SlogLevel,
		JSON:            config.JSON,
		Concise:         config.Concise,
		RequestHeaders:  config.RequestHeaders,
		ResponseHeaders: config.ResponseHeaders,
	})

	return httpLogger, nil
}
