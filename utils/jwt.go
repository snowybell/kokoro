package utils

import (
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/kelseyhightower/envconfig"
)

type JWTConfig struct {
	SecretKey string `required:"true" split_words:"true"`
}

type jwt struct {
	cfg JWTConfig
}

type JWT interface {
	New() *jwtGo.Token
	SignedString(token *jwtGo.Token) (string, error)
}

func NewJWT(cfg JWTConfig) JWT {
	return &jwt{cfg}
}

func NewJWTDefault() (JWT, error) {
	var cfg JWTConfig
	err := envconfig.Process("jwt", &cfg)
	if err != nil {
		return nil, err
	}
	return NewJWT(cfg), nil
}

func (j *jwt) New() *jwtGo.Token {
	return jwtGo.New(jwtGo.SigningMethodHS256)
}

func (j *jwt) SignedString(token *jwtGo.Token) (string, error) {
	return token.SignedString([]byte(j.cfg.SecretKey))
}
