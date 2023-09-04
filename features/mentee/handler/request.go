package handler

import "kelompok1/immersive-dash/features/mentee"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MenteeRequest struct {
	MenteeName        string `json:"mentee_name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	CurrentAdress     string `json:"current_addess"`
	HomeAdress        string `json:"home_addess"`
	Telegram          string `json:"telegram"`
	Gender            string `json:"gender"`
	EducationType     string `json:"education_type"`
	EducationMajor    string `json:"education_major"`
	EducationGraduate string `json:"education_graduate"`
	Institution       string `json:"institution"`
	EmergencyName     string `json:"emergency_name"`
	EmergencyPhone    string `json:"emergency_phone"`
	EmergencyStatus   string `json:"emergency_status"`
	StatusId          uint   `json:"status_id"`
	ClassId           uint   `json:"class_id"`
}

func RequestToCore(input MenteeRequest) mentee.CoreMentee {
	return mentee.CoreMentee{
		MenteeName:        input.MenteeName,
		Email:             input.Email,
		Phone:             input.Phone,
		CurrentAdress:     input.CurrentAdress,
		HomeAdress:        input.HomeAdress,
		Telegram:          input.Telegram,
		Gender:            input.Gender,
		EducationType:     input.EducationType,
		EducationMajor:    input.EducationMajor,
		EducationGraduate: input.EducationGraduate,
		Institution:       input.Institution,
		EmergencyName:     input.EmergencyName,
		EmergencyPhone:    input.EmergencyPhone,
		EmergencyStatus:   input.EmergencyStatus,
		StatusId:          input.StatusId,
		ClassId:           input.ClassId,
	}
}
