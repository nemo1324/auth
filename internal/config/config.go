package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	GRPC     GRPC          `envconfig:"GRPC"`
	Postgres Postgres      `envconfig:"POSTGRES"`
	TokenTTL time.Duration `envconfig:"TOKEN_TTL" default:"1h"`
}

type GRPC struct {
	Port string `envconfig:"PORT" default:":9000"`
}

type Postgres struct {
	Host     string `envconfig:"HOST" default:"localhost"`
	Port     string `envconfig:"PORT" default:"5432"`
	User     string `envconfig:"USER" default:"postgres"`
	DbName   string `envconfig:"DB_NAME" default:"postgres"`
	Password string `envconfig:"PASSWORD" default:"secret"`
	Sslmode  string `envconfig:"SSLMODE" default:"disable"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	// Загружаем переменные окружения в структуру Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to load config from environment: %v", err)
	}

	return &cfg, nil
}
