package entity

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	AccessToken  string    `gorm:"not null"`
	TokenType    string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	Expiry       time.Time `gorm:"not null"`
}
