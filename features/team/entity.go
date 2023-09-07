package team

import (
	"kelompok1/immersive-dash/features/user"
	"time"
)

type CoreTeam struct {
	ID        uint
	TeamName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Users     []user.CoreUser
}

type TeamDataInterface interface {
	SelectAll() ([]CoreTeam, error)
}

type TeamServiceInterface interface {
	GetAll() ([]CoreTeam, error)
}
