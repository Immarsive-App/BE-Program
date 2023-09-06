package handler

import "kelompok1/immersive-dash/features/mentee"

type MenteeRequest struct {
	StatusId          uint   `json:"status_id,omitempty" form:"status_id"`
	ClassId           uint   `json:"class_id,omitempty" form:"class_id"`
	FullName          string `json:"full_name,omitempty" form:"full_name"`
	CurrentAddress    string `json:"current_address,omitempty" form:"current_address"`
	HomeAddress       string `json:"home_address,omitempty" form:"home_address"`
	Email             string `json:"email,omitempty" form:"email"`
	Gender            string `json:"gender,omitempty" form:"gender"`
	Telegram          string `json:"telegram,omitempty" form:"telegram"`
	Phone             string `json:"phone,omitempty" form:"phone"`
	EmergencyName     string `json:"emergency_name,omitempty" form:"emergency_name"`
	EmergencyPhone    string `json:"emergency_phone,omitempty" form:"emergency_phone"`
	EmergencyStatus   string `json:"emergency_status,omitempty" form:"emergency_status"`
	EducationType     string `json:"education_type,omitempty" form:"education_type"`
	EducationMajor    string `json:"education_major,omitempty" form:"education_major"`
	Institution       string `json:"institution,omitempty" form:"institution"`
	EducationGraduate string `json:"education_graduate,omitempty" form:"education_graduate"`
}

func RequestToCore(input MenteeRequest) mentee.CoreMentee {
	return mentee.CoreMentee{
		FullName:          input.FullName,
		Email:             input.Email,
		Phone:             input.Phone,
		CurrentAddress:    input.CurrentAddress,
		HomeAddress:       input.HomeAddress,
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
