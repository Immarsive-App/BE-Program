package service

import "kelompok1/immersive-dash/features/feedback"

type feedbackService struct {
	feedbackData feedback.FeedbackDataInterface
}

func New(repo feedback.FeedbackDataInterface) feedback.FeedbackServiceInterface {
	return &feedbackService{
		feedbackData: repo,
	}
}
