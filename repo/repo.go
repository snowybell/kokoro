package repo

import "gorm.io/gorm"

type repository struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{db}
}
