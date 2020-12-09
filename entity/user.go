package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Username string `gorm:"unique_index;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
