package handler

import "kelompok1/immersive-dash/features/feedback"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FeedbackRequest struct {
	Note     string `json:"notes"`
	UserId   uint   `json:"user_id"`
	MenteeId uint   `json:"mentee_id"`
	StatusId uint   `json:"status_id"`
}

func RequestToCore(input FeedbackRequest) feedback.CoreFeedback {
	return feedback.CoreFeedback{
		Note:     input.Note,
		UserId:   input.UserId,
		MenteeId: input.MenteeId,
		StatusId: input.StatusId,
	}
}
