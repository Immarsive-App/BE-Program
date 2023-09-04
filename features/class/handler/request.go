package handler

import "kelompok1/immersive-dash/features/class"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClassRequest struct {
	ClassName    string `json:"class_name"`
	Mentor       string `json:"mentor"`
	StartDate    string `json:"star_date"`
	GraduateDate string `json:"graduate_date"`
}

func RequestToCore(input ClassRequest) class.CoreClass {
	return class.CoreClass{
		ClassName:   input.ClassName,
		Mentor:      input.Mentor,
		StartDate:   input.StartDate,
		GraduteDate: input.GraduateDate,
	}
}
