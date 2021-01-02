package utils

import (
	"github.com/kelseyhightower/envconfig"
)

type FrontendConfig struct {
	URL string `required:"true"`
}

func NewFrontendConfig() (*FrontendConfig, error) {
	var cfg FrontendConfig
	err := envconfig.Process("frontend", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
