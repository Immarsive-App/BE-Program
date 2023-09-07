package feedback

import (
	"time"
)

type CoreFeedback struct {
	ID        uint
	UserId    uint
	MenteeId  uint
	StatusId  uint
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserCore struct {
	ID       uint
	FullName string
}

type StatusCore struct {
	ID         uint
	StatusName string
}

type FeedbackDataInterface interface {
	SelectAll() ([]CoreFeedback, error)
	Insert(input CoreFeedback) error
	Update(id uint, input CoreFeedback) error
	Delete(id uint) error
}

type FeedbackServiceInterface interface {
	GetAll() ([]CoreFeedback, error)
	Create(input CoreFeedback) error
	Update(id uint, input CoreFeedback) error
	Deletes(id uint) error
}
