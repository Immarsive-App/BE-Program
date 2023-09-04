package user

import (
	"kelompok1/immersive-dash/features/class"
	"kelompok1/immersive-dash/features/feedback"
	"time"
)

type CoreUser struct {
	ID        uint
	FullName  string
	Email     string
	Password  string
	TeamId    uint
	Role      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Classes   []class.CoreClass
	Feedbacks []feedback.CoreFeedback
}

type UserDataInterface interface {
}

type UserServiceInterface interface {
}
