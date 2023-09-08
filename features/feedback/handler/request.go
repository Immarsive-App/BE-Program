package handler

import "kelompok1/immersive-dash/features/feedback"

type FeedbackRequest struct {
	Note     string `json:"note" form:"note"`
	UserId   uint   `json:"user_id" form:"user_id"`
	MenteeId uint   `json:"mentee_id" form:"mentee_id"`
	StatusId uint   `json:"status_id" form:"status_id"`
}

func RequestToCore(input FeedbackRequest) feedback.CoreFeedback {
	return feedback.CoreFeedback{
		Note:     input.Note,
		UserId:   input.UserId,
		MenteeId: input.MenteeId,
		StatusId: input.StatusId,
	}
}
