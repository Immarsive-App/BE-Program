package mentee

import (
	_coreClass "kelompok1/immersive-dash/features/class"
	_coreFeedback "kelompok1/immersive-dash/features/feedback"
	_coreStatus "kelompok1/immersive-dash/features/status"
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
	Feedbacks         []_coreFeedback.CoreFeedback
	Class             _coreClass.CoreClass
	Status            _coreStatus.CoreStatus
}

type MenteeDataInterface interface {
	Insert(input CoreMentee) (uint, error)
	SelectAll(className, statusName, educationType string) ([]CoreMentee, error)
	Select(menteeId uint) (CoreMentee, error)
	Update(menteeId uint, updatedMentee CoreMentee) error
	Delete(menteeId uint) error
}

type MenteeServiceInterface interface {
	Create(input CoreMentee) (uint, error)
	GetAll(className, statusName, educationType string) ([]CoreMentee, error)
	GetById(menteeId uint) (CoreMentee, error)
	Update(menteeId uint, updatedMentee CoreMentee) error
	Delete(menteeId uint) error
}
