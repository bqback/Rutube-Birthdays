package config

import (
	"birthdays/internal/apperrors"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	App         *AppConfig         `yaml:"app"`
	HttpLogging *HttpLoggingConfig `yaml:"http_logging"`
	Logging     *LoggingConfig     `yaml:"logging"`
	Database    *DatabaseConfig    `yaml:"db"`
}

type AppConfig struct {
	Port    int     `yaml:"port"`
	version float64 `yaml:"version"`
	Version string  `yaml:"-"`
}

type HttpLoggingConfig struct {
	level           string     `yaml:"level"`
	Level           slog.Level `yaml:"-"`
	JSON            bool       `yaml:"json_logs"`
	Concise         bool       `yaml:"concise_logs"`
	RequestHeaders  bool       `yaml:"include_request_headers"`
	ResponseHeaders bool       `yaml:"include_response_headers"`
}

type LoggingConfig struct {
	Level                  string `yaml:"level"`
	DisableTimestamp       bool   `yaml:"disable_timestamp"`
	FullTimestamp          bool   `yaml:"full_timestamp"`
	DisableLevelTruncation bool   `yaml:"disable_level_truncation"`
	LevelBasedReport       bool   `yaml:"level_based_report"`
	ReportCaller           bool   `yaml:"report_caller"`
}

type DatabaseConfig struct {
	User              string `yaml:"user"`
	Password          string `yaml:"-"`
	Host              string `yaml:"-"`
	Port              uint64 `yaml:"port"`
	DBName            string `yaml:"db_name"`
	AppName           string `yaml:"app_name"`
	Schema            string `yaml:"schema"`
	ConnectionTimeout uint64 `yaml:"connection_timeout"`
}

var loggingLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func LoadConfig(configPath, envPath string) (*Config, error) {
	var (
		config Config
		err    error
	)

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	if envPath == "" {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(envPath)
	}

	if err != nil {
		return nil, apperrors.ErrEnvNotFound
	}

	config.Database.Password, err = GetDBPassword()
	if err != nil {
		return nil, err
	}

	config.Database.Host = GetDBConnectionHost()

	level, ok := loggingLevels[config.HttpLogging.level]
	if !ok {
		return nil, apperrors.ErrInvalidLoggingLevel
	}
	config.HttpLogging.Level = level

	config.App.Version = fmt.Sprintf("v%f", config.App.version)

	return &config, err
}

// GetDBConnectionHost
// возвращает имя хоста из env для соединения с БД (по умолчанию localhost)
func GetDBConnectionHost() string {
	host, hOk := os.LookupEnv("POSTGRES_HOST")
	if !hOk {
		return "localhost"
	}
	return host
}

// getDBConnectionHost
// возвращает пароль из env для соединения с БД
func GetDBPassword() (string, error) {
	pwd, pOk := os.LookupEnv("POSTGRES_PASSWORD")
	if !pOk {
		return "", apperrors.ErrDatabasePWMissing
	}
	return pwd, nil
}
