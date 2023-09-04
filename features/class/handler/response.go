package handler

type ClassResponse struct {
	ID           uint   `json:"id"`
	ClassName    string `json:"class_name"`
	Mentor       string `json:"mentor"`
	StartDate    string `json:"star_date"`
	GraduateDate string `json:"graduate_date"`
}
