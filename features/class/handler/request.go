package handler

import "kelompok1/immersive-dash/features/class"

type ClassRequest struct {
	ClassName    string `json:"class_name"`
	Mentor       string `json:"mentor"`
	StartDate    string `json:"star_date"`
	GraduateDate string `json:"graduate_date"`
}

func RequestToCore(input ClassRequest) class.CoreClass {
	return class.CoreClass{
		ClassName: input.ClassName,
	}
}
