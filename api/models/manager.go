package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Website string `gorm:"not null" json:"website"`
	WPass   string `gorm:"not null" json:"wpass"`
	UserID  int    `gorm:"not null" json:"user_id"`
}
