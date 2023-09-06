package status

import (
	"kelompok1/immersive-dash/features/feedback"
	"kelompok1/immersive-dash/features/mentee"
	"time"
)

type CoreStatus struct {
	ID         uint
	StatusName string `gorm:"status_name;unique;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	Mentees    []mentee.CoreMentee
	Feedbacks  []feedback.CoreFeedback
}

type StatusDataInterface interface {
	SelectAll() ([]CoreStatus, error)
}

type StatusServiceInterface interface {
	GetAll() ([]CoreStatus, error)
}
