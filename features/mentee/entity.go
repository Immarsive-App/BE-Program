package mentee

import (
	"kelompok1/immersive-dash/features/feedback"
	"time"
)

type CoreMentee struct {
	ID                uint
	StatusId          uint
	ClassId           uint
	MenteeName        string
	CurrentAddress    string
	HomeAddress       string
	Email             string
	Gender            string
	Telegram          string
	Phone             string
	EmergencyName     string
	EmergencyPhone    string
	EmergencyStatus   string
	EducationType     string
	EducationMajor    string
	Institution       string
	EducationGraduate string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
	Feedbacks       []feedback.CoreFeedback
}

type MenteeDataInterface interface {
	// SelectAll() ([]Core, error)
}

type MenteeServiceInterface interface {
	// GetAll() ([]Core, error)
}
