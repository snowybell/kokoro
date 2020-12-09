package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *Config) (*gorm.DB, error) {
	dsn := postgres.Open(cfg.ToConnString())
	return gorm.Open(dsn, &gorm.Config{
		AllowGlobalUpdate: false,
	})
}

func NewDBDefault() (*gorm.DB, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	return NewDB(cfg)
}
