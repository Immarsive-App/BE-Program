package service

import (
	"errors"
	"kelompok1/immersive-dash/features/feedback"
)

type feedbackService struct {
	feedbackData feedback.FeedbackDataInterface
}

// GetAll implements feedback.FeedbackServiceInterface
func (service *feedbackService) GetAll() ([]feedback.CoreFeedback, error) {
	result, err := service.feedbackData.SelectAll()
	return result, err
}

// Deletes implements feedback.FeedbackServiceInterface
func (service *feedbackService) Deletes(id uint) error {
	err := service.feedbackData.Delete(id)
	return err
}

// Update implements feedback.FeedbackServiceInterface
func (service *feedbackService) Update(id uint, input feedback.CoreFeedback) error {
	err := service.feedbackData.Update(id, input)
	return err
}

// Create implements feedback.FeedbackServiceInterface
func (service *feedbackService) Create(input feedback.CoreFeedback) error {
	if input.Note == "" {
		return errors.New("validation error. name/detail/deskiption required")
	}
	err := service.feedbackData.Insert(input)
	return err
}

func New(repo feedback.FeedbackDataInterface) feedback.FeedbackServiceInterface {
	return &feedbackService{
		feedbackData: repo,
	}
}
