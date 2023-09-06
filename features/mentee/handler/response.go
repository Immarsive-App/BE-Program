package handler

import (
	"kelompok1/immersive-dash/features/mentee"
)

type CreateDeleteResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Telegram string `json:"telegram"`
	Gender   string `json:"gender"`
}
type GetAllMenteeResponse struct {
	ID            uint   `json:"id"`
	ClassID       uint   `json:"class_id"`
	FullName      string `json:"full_name"`
	StatusID      uint   `json:"status_id"`
	EducationType string `json:"education_type"`
	Gender        string `json:"gender"`
}
type UpdateResponse struct {
	ID                uint   `json:"mentee_id,omitempty" form:"mentee_id"`
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
	Institution       string `json:"institution,omitempty" form:"institution"`
	EducationMajor    string `json:"major,omitempty" form:"major"`
	EducationGraduate string `json:"graduate,omitempty" form:"graduate"`
}

type GetByIdResponse struct {
	ID                uint   `json:"mentee_id,omitempty" form:"mentee_id"`
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
	Institution       string `json:"institution,omitempty" form:"institution"`
	EducationMajor    string `json:"major,omitempty" form:"major"`
	EducationGraduate string `json:"graduate,omitempty" form:"graduate"`
}

func CoreToCreateDeleteResponse(mentee mentee.CoreMentee) CreateDeleteResponse {
	return CreateDeleteResponse{
		ID:       mentee.ID,
		FullName: mentee.FullName,
		Email:    mentee.Email,
		Phone:    mentee.Phone,
		Telegram: mentee.Telegram,
		Gender:   mentee.Gender,
	}
}

func CoreToUpdateResponse(mentee mentee.CoreMentee) UpdateResponse {
	return UpdateResponse{
		ID:                mentee.ID,
		StatusId:          mentee.StatusId,
		ClassId:           mentee.ClassId,
		FullName:          mentee.FullName,
		CurrentAddress:    mentee.CurrentAddress,
		HomeAddress:       mentee.HomeAddress,
		Email:             mentee.Email,
		Gender:            mentee.Gender,
		Telegram:          mentee.Telegram,
		Phone:             mentee.Phone,
		EmergencyName:     mentee.EmergencyName,
		EmergencyPhone:    mentee.EmergencyPhone,
		EmergencyStatus:   mentee.EmergencyStatus,
		EducationType:     mentee.EducationType,
		Institution:       mentee.Institution,
		EducationMajor:    mentee.EducationMajor,
		EducationGraduate: mentee.EducationGraduate,
	}
}

func CoreToGetByIdResponse(mentee mentee.CoreMentee) GetByIdResponse {
	return GetByIdResponse{
		ID:                mentee.ID,
		StatusId:          mentee.StatusId,
		ClassId:           mentee.ClassId,
		FullName:          mentee.FullName,
		CurrentAddress:    mentee.CurrentAddress,
		HomeAddress:       mentee.HomeAddress,
		Email:             mentee.Email,
		Gender:            mentee.Gender,
		Telegram:          mentee.Telegram,
		Phone:             mentee.Phone,
		EmergencyName:     mentee.EmergencyName,
		EmergencyPhone:    mentee.EmergencyPhone,
		EmergencyStatus:   mentee.EmergencyStatus,
		EducationType:     mentee.EducationType,
		Institution:       mentee.Institution,
		EducationMajor:    mentee.EducationMajor,
		EducationGraduate: mentee.EducationGraduate,
	}
}

func CoreToGetAllResponse(mentee mentee.CoreMentee) GetAllMenteeResponse {
	return GetAllMenteeResponse{
		ID:            mentee.ID,
		ClassID:       mentee.ClassId,
		FullName:      mentee.FullName,
		StatusID:      mentee.StatusId,
		EducationType: mentee.EducationType,
		Gender:        mentee.Gender,
	}
}
