package handler

type FeedbackResponse struct {
	ID       uint   `json:"id"`
	Note     string `json:"notes"`
	UserId   uint   `json:"user_id"` //user yang kasih feedback dari user id ke nama user
	StatusId uint   `json:"status_id"`
	FullName string `json:"user"`
	Status   string `json:"status"`
}
