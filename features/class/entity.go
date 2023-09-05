package class

import (
	"kelompok1/immersive-dash/features/mentee"
	"time"
)

type CoreClass struct {
	ID        uint
	UserId    uint
	ClassName string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Mentee    []mentee.CoreMentee
}

type ClassDataInterface interface {
	SelectAll() ([]CoreClass, error)
	Insert(input CoreClass) error
	SelectById(id uint) (CoreClass, error)
	Update(id uint, input CoreClass) error
	Delete(id uint) error
}

type ClassServiceInterface interface {
	GetAll() ([]CoreClass, error)
	Create(input CoreClass) error
	GetById(id uint) (CoreClass, error)
	Update(id uint, input CoreClass) error
	Deletes(id uint) error
}
