package handler

import "kelompok1/immersive-dash/features/feedback"

type FeedbackResponse struct {
	ID       uint   `json:"id"`
	Note     string `json:"notes"`
	UserId   uint   `json:"user_id"` //user yang kasih feedback dari user id ke nama user
	FullName string `json:"user"`
	Status   string `json:"status"`
}

func CoreToResponse(core feedback.CoreFeedback) FeedbackResponse {
	return FeedbackResponse{
		ID:     core.ID,
		Note:   core.Note,
		UserId: core.UserId,
	}
}
