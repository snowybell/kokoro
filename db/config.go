package db

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string `required:"true"`
	Name     string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	SSLMode  string `default:"disable"`
}

func (cfg *Config) ToConnString() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Name, cfg.Name, cfg.SSLMode, cfg.Password)
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("database", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
