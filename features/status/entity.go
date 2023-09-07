package status

import (
	"kelompok1/immersive-dash/features/feedback"
	"time"
)

type CoreStatus struct {
	ID         uint
	StatusName string `gorm:"status_name;unique;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	Feedbacks  []feedback.CoreFeedback
}

type StatusDataInterface interface {
	// SelectAll() ([]Core, error)
}

type StatusServiceInterface interface {
	// GetAll() ([]Core, error)
}
