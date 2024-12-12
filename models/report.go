package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Category    string  `json:"category" gorm:"not null"`
	Title       string  `json:"title" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
	PositionX   float64 `json:"positionX" gorm:"not null"`
	PositionY   float64 `json:"positionY" gorm:"not null"`
}