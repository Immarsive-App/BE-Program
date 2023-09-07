package handler

import "kelompok1/immersive-dash/features/mentee"

type MenteeRequest struct {
	StatusId          uint   `json:"status_Id" form:"status"`
	ClassId           uint   `json:"class_id" form:"class_id"`
	FullName          string `json:"full_name" form:"full_name"`
	CurrentAddress    string `json:"current_address" form:"current_address"`
	HomeAddress       string `json:"home_address" form:"home_address"`
	Email             string `json:"email" form:"email"`
	Gender            string `json:"gender" form:"gender"`
	Telegram          string `json:"telegram" form:"telegram"`
	Phone             string `json:"phone" form:"phone"`
	EmergencyName     string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone    string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus   string `json:"emergency_status" form:"emergency_status"`
	EducationType     string `json:"education_type" form:"education_type"`
	EducationMajor    string `json:"education_major" form:"education_major"`
	Institution       string `json:"institution" form:"institution"`
	EducationGraduate string `json:"education_graduate" form:"education_graduate"`
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

type UpdateMenteeRequest struct {
	ClassId  uint   `json:"class_id" form:"class_id"`
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Gender   string `json:"gender" form:"gender"`
	Telegram string `json:"telegram" form:"telegram"`
	Phone    string `json:"phone" form:"phone"`
}

func UpdateRequestToCore(input UpdateMenteeRequest) mentee.CoreMentee {
	return mentee.CoreMentee{
		FullName: input.FullName,
		Email:    input.Email,
		Phone:    input.Phone,
		Telegram: input.Telegram,
		Gender:   input.Gender,
		ClassId:  input.ClassId,
	}
}
