package class

import (
	"time"
)

type CoreClass struct {
	ID        uint
	UserId    uint
	ClassName string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
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
