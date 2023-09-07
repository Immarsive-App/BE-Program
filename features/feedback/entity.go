package feedback

import "time"

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

type FeedbackDataInterface interface {
	// SelectAll() ([]Core, error)
}

type FeedbackServiceInterface interface {
	// GetAll() ([]Core, error)
}
