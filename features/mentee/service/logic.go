package service

import (
	"fmt"
	"kelompok1/immersive-dash/features/mentee"

	"github.com/go-playground/validator/v10"
)

type menteeService struct {
	menteeData mentee.MenteeDataInterface
	validate   *validator.Validate
}

func New(repo mentee.MenteeDataInterface) mentee.MenteeServiceInterface {
	return &menteeService{
		menteeData: repo,
		validate:   validator.New(),
	}
}

// Create implements mentee.MenteeServiceInterface.
func (service *menteeService) Create(input mentee.CoreMentee) (uint, error) {
	validationErr := service.validate.Struct(input)
	if validationErr != nil {
		return 0, validationErr
	}
	menteeId, err := service.menteeData.Insert(input)
	if err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}
	return menteeId, nil
}

// GetAll implements mentee.MenteeServiceInterface.
func (service *menteeService) GetAll(className, statusName, educationType string) ([]mentee.CoreMentee, error) {

	result, err := service.menteeData.SelectAll(className, statusName, educationType)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return result, nil
}

// GetById implements mentee.MenteeServiceInterface.
func (service *menteeService) GetById(menteeId uint) (mentee.CoreMentee, error) {
	result, err := service.menteeData.Select(menteeId)
	if err != nil {
		return mentee.CoreMentee{}, fmt.Errorf("error: %v", err)
	}
	return result, nil
}

// Delete implements mentee.MenteeServiceInterface.
func (service *menteeService) Delete(menteeId uint) error {
	err := service.menteeData.Delete(menteeId)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}

// Update implements mentee.MenteeServiceInterface.
func (service *menteeService) Update(menteeId uint, updatedMentee mentee.CoreMentee) error {
	err := service.menteeData.Update(menteeId, updatedMentee)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
