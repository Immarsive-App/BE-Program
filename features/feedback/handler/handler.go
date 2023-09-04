package handler

import (
	"kelompok1/immersive-dash/features/feedback"
)

type FeedbackHandler struct {
	feedbackService feedback.FeedbackServiceInterface
}

func New(service feedback.FeedbackServiceInterface) *FeedbackHandler {
	return &FeedbackHandler{
		feedbackService: service,
	}
}
