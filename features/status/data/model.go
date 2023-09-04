package data

import (
	_feedbackData "kelompok1/immersive-dash/features/feedback/data"
	_menteeData "kelompok1/immersive-dash/features/mentee/data"

	"gorm.io/gorm"
)

// struct Status gorm model
type Status struct {
	gorm.Model
	StatusName string                   `gorm:"status_name;unique;not null"`
	Mentees    []_menteeData.Mentee     `gorm:"foreignKey:StatusId"`
	Feedbacks  []_feedbackData.Feedback `gorm:"foreignKey:StatusId"`
}
