package repo

import "github.com/snowybell/kokoro/entity"

type Repository interface {
	// User methods
	GetUser(user entity.User) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	SaveUser(user entity.User) (*entity.User, error)
	DeleteUser(user entity.User) error

	// Token methods
	GetToken(token entity.Token) (*entity.Token, error)
	SaveToken(token entity.Token) (*entity.Token, error)
	DeleteToken(token entity.Token) error
}
