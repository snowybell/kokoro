package entity

import "golang.org/x/oauth2"

type Passport struct {
	User       *User         `json:"user"`
	OAuthToken *oauth2.Token `json:"oauth_token"`
}
