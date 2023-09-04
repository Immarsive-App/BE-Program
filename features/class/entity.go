package class

import (
	"kelompok1/immersive-dash/features/mentee"
	"time"
)

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
	Mentee      []mentee.CoreMentee
}

type ClassDataInterface interface {
}

type ClassServiceInterface interface {
}
