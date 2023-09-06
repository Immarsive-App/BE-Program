package mentee

import (
	"kelompok1/immersive-dash/features/feedback"
	"time"
)

type CoreMentee struct {
	ID                uint
	StatusId          uint
	ClassId           uint
	FullName          string
	CurrentAddress    string
	HomeAddress       string
	Email             string /*`validate:"required,email"`*/
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
	Feedbacks         []feedback.CoreFeedback
}

type MenteeDataInterface interface {
	Insert(input CoreMentee) (uint, error)
	SelectAll() ([]CoreMentee, error)
	Select(menteeId uint) (CoreMentee, error)
	Update(menteeId uint, updatedMentee CoreMentee) error
	Delete(menteeId uint) error
}

type MenteeServiceInterface interface {
	Create(input CoreMentee) (uint, error)
	GetAll() ([]CoreMentee, error)
	GetById(menteeId uint) (CoreMentee, error)
	Update(menteeId uint, updatedMentee CoreMentee) error
	Delete(menteeId uint) error
}
