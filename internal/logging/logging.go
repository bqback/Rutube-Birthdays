package logging

import (
	"birthdays/internal/config"

	"github.com/go-chi/httplog/v2"
	"github.com/sirupsen/logrus"
)

type ILogger interface {
	DebugFmt(message string, request string, function string, routeNode string)
	DebugRequestlessFmt(message string, function string, routeNode string)
	Debug(message string)
	Info(message string)
	Error(message string)
	Fatal(message string)
	Printf(message string, args ...interface{})
	Level() int
}

type LogrusLogger struct {
	level  int
	logger *logrus.Logger
}

type Loggers struct {
	HTTP   *httplog.Logger
	Logrus *LogrusLogger
}

func Setup(config *config.Config) (*Loggers, error) {
	httpLogger := httplog.NewLogger("birthdays-api", httplog.Options{
		LogLevel:        config.HttpLogging.Level,
		JSON:            config.HttpLogging.JSON,
		Concise:         config.HttpLogging.Concise,
		RequestHeaders:  config.HttpLogging.RequestHeaders,
		ResponseHeaders: config.HttpLogging.ResponseHeaders,
	})

	logrusLogger, err := NewLogrusLogger(config.Logging)
	if err != nil {
		return nil, err
	}

	loggers := Loggers{
		HTTP:   httpLogger,
		Logrus: &logrusLogger,
	}

	return &loggers, nil
}
