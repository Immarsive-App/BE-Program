package handler

type FeedbackResponse struct {
	ID       uint   `json:"id"`
	MenteeId uint   `json:"mentee_id"`
	Note     string `json:"note"`
	UserId   uint   `json:"user_id"` //user yang kasih feedback dari user id ke nama user
	StatusId uint   `json:"status"`
}
