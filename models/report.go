package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Category    string  `gorm:"not null"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	PositionX   float32 `gorm:"not null"`
	PositionY   float32 `gorm:"not null"`
}
