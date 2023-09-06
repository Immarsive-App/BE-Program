package mentee

import (
	"kelompok1/immersive-dash/features/feedback"
	"time"
)

type CoreMentee struct {
	ID                uint
	StatusId          uint `validate:"required"`
	ClassId           uint
	FullName          string `validate:"required"`
	CurrentAddress    string `validate:"required"`
	HomeAddress       string `validate:"required"`
	Email             string `validate:"required,email"`
	Gender            string `validate:"required"`
	Telegram          string `validate:"required"`
	Phone             string `validate:"required"`
	EmergencyName     string `validate:"required"`
	EmergencyPhone    string `validate:"required"`
	EmergencyStatus   string `validate:"required"`
	EducationType     string `validate:"required"`
	EducationMajor    string `validate:"required"`
	Institution       string `validate:"required"`
	EducationGraduate string `validate:"required"`
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
