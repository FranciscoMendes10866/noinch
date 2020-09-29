package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"not null" json:"id"`
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
