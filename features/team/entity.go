package team

import "time"

type CoreTeam struct {
	ID        uint
	TeamName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type TeamDataInterface interface {
	// SelectAll() ([]Core, error)
}

type TeamServiceInterface interface {
	// GetAll() ([]Core, error)
}
