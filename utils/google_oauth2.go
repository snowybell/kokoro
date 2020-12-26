package utils

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuthConfig struct {
	ClientID     string `required:"true" split_words:"true"`
	ClientSecret string `required:"true" split_words:"true"`
	RedirectURL  string `required:"true" split_words:"true"`
}

func NewGoogleOAuthConfig() (*oauth2.Config, error) {
	var cfg GoogleOAuthConfig
	err := envconfig.Process("GOOGLE_OAUTH2", &cfg)
	if err != nil {
		return nil, err
	}
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/drive.appdata",
			"https://www.googleapis.com/auth/drive.file",
		},
		Endpoint: google.Endpoint,
	}, nil
}
