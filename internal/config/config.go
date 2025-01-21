package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	GRPC     GRPC          `yaml:"grpc"`
	Postgres Postgres      `yaml:"postgres"`
	TokenTTL time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GRPC struct {
	Port string `yaml:"port"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DbName   string `yaml:"dbName"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	// Открываем YAML файл
	file, err := os.Open("config/local.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	// Декодируем содержимое YAML-файла
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	return &cfg, nil
}
