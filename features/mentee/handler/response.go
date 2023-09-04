package handler

type MenteeResponse struct {
	ID         uint   `json:"id"`
	MenteeName string `json:"mentee_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Telegram   string `json:"telegram"`
}
