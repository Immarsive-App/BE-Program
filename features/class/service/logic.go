package service

import (
	"errors"
	"kelompok1/immersive-dash/features/class"
)

type classService struct {
	classData class.ClassDataInterface
}

// Deletes implements class.ClassServiceInterface
func (service *classService) Deletes(id uint) error {
	err := service.classData.Delete(id)
	return err
}

// Update implements class.ClassServiceInterface
func (service *classService) Update(id uint, input class.CoreClass) error {
	err := service.classData.Update(id, input)
	return err
}

// GetById implements class.ClassServiceInterface
func (repo *classService) GetById(id uint) (class.CoreClass, error) {
	return repo.classData.SelectById(id)
}

// Create implements class.ClassServiceInterface
func (service *classService) Create(input class.CoreClass) error {
	if input.ClassName == "" {
		return errors.New("validation error. class_name required")
	}
	err := service.classData.Insert(input)
	return err
}

// GetAll implements class.ClassServiceInterface
func (service *classService) GetAll() ([]class.CoreClass, error) {
	result, err := service.classData.SelectAll()
	return result, err
}

func New(repo class.ClassDataInterface) class.ClassServiceInterface {
	return &classService{
		classData: repo,
	}
}
