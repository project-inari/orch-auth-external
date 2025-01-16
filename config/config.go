// Package config provides configuration settings for the server
package config

import (
	"log"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var once sync.Once
var config *Config

// New loads the configuration from the .env file
func New() *Config {
	once.Do(func() {
		e := os.Getenv("APP_ENV_STAGE")
		if e == "" || e == "LOCAL" {
			if err := godotenv.Load(".env.generated"); err != nil {
				slog.Warn("[config.New] unable to load .env.generated file", slog.Any("error", err))
			}
		}

		cfg := &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Panicf("error - [config.New] unable to parse config: %v", err)
		}
		config = cfg
	})

	return config
}

// Config represents the configuration of the server
type Config struct {
	AppConfig         AppConfig
	LogConfig         LogConfig
	SentryConfig      SentryConfig
	CoreAuthAPIConfig CoreAuthAPIConfig
	CoreUserAPIConfig CoreUserAPIConfig
}

// AppConfig represents the configuration of the application
type AppConfig struct {
	Name     string `env:"APP_NAME,notEmpty"`
	Port     string `env:"APP_PORT,notEmpty"`
	EnvStage string `env:"APP_ENV_STAGE,notEmpty"`
}

// LogConfig represents the configuration of the logger
type LogConfig struct {
	Level             string `env:"LOG_LEVEL,notEmpty"`
	MaskSensitiveData bool   `env:"LOG_MASK_SENSITIVE_DATA,notEmpty"`
}

// SentryConfig represents the configuration of Sentry.io
type SentryConfig struct {
	SentryDSN string `env:"SENTRY_DSN"`
}

// CoreAuthAPIConfig represents the configuration of the core auth API
type CoreAuthAPIConfig struct {
	BaseURL                  string        `env:"CORE_AUTH_API_BASE_URL,notEmpty"`
	SignupPath               string        `env:"CORE_AUTH_API_SIGNUP_PATH,notEmpty"`
	DeleteUserPath           string        `env:"CORE_AUTH_API_DELETE_USER_PATH,notEmpty"`
	MaxConns                 int           `env:"CORE_AUTH_API_MAX_CONNS,notEmpty"`
	MaxRetry                 int           `env:"CORE_AUTH_API_MAX_RETRY,notEmpty"`
	Timeout                  time.Duration `env:"CORE_AUTH_API_TIMEOUT,notEmpty"`
	InsecureSkipVerify       bool          `env:"CORE_AUTH_API_INSECURE_SKIP_VERIFY,notEmpty"`
	MaxTransactionsPerSecond int           `env:"CORE_AUTH_API_MAX_TRANSACTIONS_PER_SECOND"`
}

// CoreUserAPIConfig represents the configuration of the core user API
type CoreUserAPIConfig struct {
	BaseURL                  string        `env:"CORE_USER_API_BASE_URL,notEmpty"`
	SignupPath               string        `env:"CORE_USER_API_SIGNUP_PATH,notEmpty"`
	MaxConns                 int           `env:"CORE_USER_API_MAX_CONNS,notEmpty"`
	MaxRetry                 int           `env:"CORE_USER_API_MAX_RETRY,notEmpty"`
	Timeout                  time.Duration `env:"CORE_USER_API_TIMEOUT,notEmpty"`
	InsecureSkipVerify       bool          `env:"CORE_USER_API_INSECURE_SKIP_VERIFY,notEmpty"`
	MaxTransactionsPerSecond int           `env:"CORE_USER_API_MAX_TRANSACTIONS_PER_SECOND"`
}
