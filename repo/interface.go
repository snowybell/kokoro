package repo

import "github.com/snowybell/kokoro/entity"

type Repository interface {
	// User methods
	GetUser(user entity.User) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	SaveUser(user entity.User) (*entity.User, error)
	DeleteUser(user entity.User) error
}
