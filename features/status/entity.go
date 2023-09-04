package status

import "time"

type CoreStatus struct {
	ID         uint
	StatusName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type StatusDataInterface interface {
	// SelectAll() ([]Core, error)
}

type StatusServiceInterface interface {
	// GetAll() ([]Core, error)
}
