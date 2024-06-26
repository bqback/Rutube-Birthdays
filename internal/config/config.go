package config

import (
	"birthdays/internal/apperrors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	yaml "gopkg.in/yaml.v3"
)

var loggingLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

type Config struct {
	App      *AppConfig      `yaml:"app"`
	Logging  *LoggingConfig  `yaml:"logging"`
	Database *DatabaseConfig `yaml:"db"`
	JWT      *JWTConfig      `yaml:"jwt"`
}

type AppConfig struct {
	Port    int     `yaml:"port"`
	version float64 `yaml:"version"`
	Version string  `yaml:"-"`
}

type LoggingConfig struct {
	Level           string     `yaml:"level"`
	SlogLevel       slog.Level `yaml:"-"`
	JSON            bool       `yaml:"json_logs"`
	Concise         bool       `yaml:"concise_logs"`
	RequestHeaders  bool       `yaml:"include_request_headers"`
	ResponseHeaders bool       `yaml:"include_response_headers"`
}

func (c *LoggingConfig) SetLevel() error {
	slogLevel, ok := loggingLevels[c.Level]
	if !ok {
		log.Println("level in config", c.Level)
		return apperrors.ErrInvalidLoggingLevel
	}
	c.SlogLevel = slogLevel
	return nil
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

type JWTConfig struct {
	Secret          string        `yaml:"-"`
	LifetimeSeconds uint          `yaml:"lifetime_seconds"`
	Lifetime        time.Duration `yaml:"-"`
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

	config.Database.Password, err = getDBPassword()
	if err != nil {
		return nil, err
	}

	config.Database.Host = getDBConnectionHost()

	err = config.Logging.SetLevel()
	if err != nil {
		return nil, err
	}

	config.JWT.Secret, err = getJWTSecret()
	if err != nil {
		return nil, err
	}
	config.JWT.Lifetime = time.Duration(config.JWT.LifetimeSeconds) * time.Second

	config.App.Version = fmt.Sprintf("v%f", config.App.version)

	return &config, err
}

// getDBConnectionHost
// возвращает имя хоста из env для соединения с БД (по умолчанию localhost)
func getDBConnectionHost() string {
	host, hOk := os.LookupEnv("POSTGRES_HOST")
	if !hOk {
		return "localhost"
	}
	return host
}

// getDBConnectionHost
// возвращает пароль из env для соединения с БД
func getDBPassword() (string, error) {
	pwd, pOk := os.LookupEnv("POSTGRES_PASSWORD")
	if !pOk {
		return "", apperrors.ErrDatabasePWMissing
	}
	return pwd, nil
}

func getJWTSecret() (string, error) {
	name, nOk := os.LookupEnv("JWT_SECRET")
	if !nOk {
		return "", apperrors.ErrJWTSecretMissing
	}
	return name, nil
}
