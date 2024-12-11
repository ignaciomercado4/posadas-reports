package models

type ReportInput struct {
	Category    string  `gorm:"not null"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	PositionX   float32 `gorm:"not null"`
	PositionY   float32 `gorm:"not null"`
}
