package service

import (
	"errors"
	"kelompok1/immersive-dash/features/feedback"
)

type feedbackService struct {
	feedbackData feedback.FeedbackDataInterface
}

// Create implements feedback.FeedbackServiceInterface
func (service *feedbackService) Create(input feedback.CoreFeedback, userId uint) (feedback.CoreFeedback, error) {
	if input.Note == "" || input.MenteeId == 0 || input.StatusId == 0 {
		return feedback.CoreFeedback{}, errors.New("validation error. Note tidak boleh kosong, meentee_id dan, status_id harus sesuai")
	}
	result, err := service.feedbackData.Insert(input, userId)
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

func New(repo feedback.FeedbackDataInterface) feedback.FeedbackServiceInterface {
	return &feedbackService{
		feedbackData: repo,
	}
}
