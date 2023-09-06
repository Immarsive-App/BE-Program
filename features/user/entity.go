package user

import (
	"kelompok1/immersive-dash/features/class"
	"kelompok1/immersive-dash/features/feedback"
	"time"
)

type CoreUser struct {
	ID       uint
	FullName string
	Email    string
	Password string
	TeamId   uint
	Role     string
	//Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Classes   []class.CoreClass
	Feedbacks []feedback.CoreFeedback
}

type UserDataInterface interface {
	SelectAll() ([]CoreUser, error)
	Insert(input CoreUser) error
	Login(email string) (CoreUser, error)
	Delete(id uint) error
	GetByID(id uint) (CoreUser, error)
	Update(id uint, updatedUser CoreUser) error
}

type UserServiceInterface interface {
	GetAll() ([]CoreUser, error)
	Create(input CoreUser) error
	Login(email string, password string) (CoreUser, string, error)
	GetByID(id uint) (CoreUser, error)
	Delete(id uint) error
	Update(id uint, updatedUser CoreUser) error
}
