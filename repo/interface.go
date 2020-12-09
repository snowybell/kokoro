package repo

import "github.com/snowybell/kokoro/entity"

type Repository interface {
	// User methods
	SaveUser(user entity.User) error
	GetUser(user entity.User) (*entity.User, error)
	DeleteUser(user entity.User) error
}
