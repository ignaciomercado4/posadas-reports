package models

type ReportInput struct {
	Category    string  `json:"category" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Urgency     string  `json:"urgency" binding:"required"`
	Address     string  `json:"address" gorm:"not null"`
	PositionX   float64 `json:"positionX" binding:"required"`
	PositionY   float64 `json:"positionY" binding:"required"`
}
