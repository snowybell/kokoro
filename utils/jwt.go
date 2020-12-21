package utils

import (
	"github.com/kelseyhightower/envconfig"
)

type JWTConfig struct {
	SecretKey []byte `required:"true" split_words:"true"`
}

func NewJWTConfig() (*JWTConfig, error) {
	var cfg JWTConfig
	err := envconfig.Process("jwt", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
