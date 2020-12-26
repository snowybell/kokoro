package db

import (
	"github.com/snowybell/kokoro/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *Config) (*gorm.DB, error) {
	dsn := postgres.Open(cfg.ToConnString())
	db, err := gorm.Open(dsn, &gorm.Config{
		AllowGlobalUpdate: false,
	})

	if err != nil {
		return nil, err
	}

	if cfg.AutoMigrate {
		err = AutoMigrate(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func NewDBDefault() (*gorm.DB, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	return NewDB(cfg)
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Token{},
	)
}
