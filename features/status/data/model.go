package data

import (
	_feedbackData "kelompok1/immersive-dash/features/feedback/data"

	"gorm.io/gorm"
)

// struct Status gorm model
type Status struct {
	gorm.Model
	StatusName string                   `gorm:"status_name;unique;not null"`
	Feedbacks  []_feedbackData.Feedback `gorm:"foreignKey:StatusId"`
}
