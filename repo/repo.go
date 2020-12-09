package repo

import (
	"github.com/snowybell/kokoro/db"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{db}
}

func NewRepoDefault() (Repository, error) {
	dbConnection, err := db.NewDBDefault()
	if err != nil {
		return nil, err
	}
	return NewRepo(dbConnection), nil
}
