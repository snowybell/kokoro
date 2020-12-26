package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `gorm:"not null" json:"name"`
	Email   string `gorm:"uniqueIndex" json:"email"`
	TokenID uint   `gorm:"not null"`
}
