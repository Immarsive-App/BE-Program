package status

import "time"

type CoreStatus struct {
	ID         uint
	StatusName string `gorm:"status_name;unique;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type StatusDataInterface interface {
	// SelectAll() ([]Core, error)
}

type StatusServiceInterface interface {
	// GetAll() ([]Core, error)
}
