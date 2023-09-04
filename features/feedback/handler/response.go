package handler

type FeedbackResponse struct {
	ID       uint   `json:"id"`
	Note     string `json:"notes"`
	FullName string `json:"full_name"`
	Status   string `json:"status"`
}
