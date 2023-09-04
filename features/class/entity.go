package class

import "time"

type CoreClass struct {
	ID          uint
	UserId      uint
	ClassName   string
	Mentor      string
	StartDate   string
	GraduteDate string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type ClassDataInterface interface {
	// SelectAll() ([]Core, error)
}

type ClassServiceInterface interface {
	// GetAll() ([]Core, error)
}
