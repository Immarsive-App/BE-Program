package handler

import (
	"kelompok1/immersive-dash/features/feedback"
	"kelompok1/immersive-dash/features/mentee"
)

type CreatDeleteMenteeResponse struct {
	ID       uint   `json:"id,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Telegram string `json:"telegram,omitempty"`
	Gender   string `json:"gender,omitempty"`
}
type GetAllMenteeResponse struct {
	ID            uint   `json:"id,omitempty"`
	FullName      string `json:"full_name,omitempty"`
	Class         string `json:"class,omitempty"`
	Status        string `json:"status,omitempty"`
	EducationType string `json:"education_type,omitempty"`
	Gender        string `json:"gender,omitempty"`
}
type GetUpdateMenteeResponse struct {
	ID                uint   `json:"id,omitempty"`
	Class             string `json:"class,omitempty"`
	FullName          string `json:"full_name,omitempty"`
	CurrentAddress    string `json:"current_address,omitempty"`
	HomeAddress       string `json:"home_address,omitempty"`
	Email             string `json:"email,omitempty"`
	Gender            string `json:"gender,omitempty"`
	Telegram          string `json:"telegram,omitempty"`
	Phone             string `json:"phone,omitempty" `
	EmergencyName     string `json:"emergency_name,omitempty"`
	EmergencyPhone    string `json:"emergency_phone,omitempty"`
	EmergencyStatus   string `json:"emergency_status,omitempty"`
	EducationType     string `json:"education_type,omitempty"`
	Institution       string `json:"institution,omitempty"`
	EducationMajor    string `json:"major,omitempty"`
	EducationGraduate string `json:"graduate,omitempty"`
}

type GetMenteeFeedbackResponse struct {
	ID             uint   `json:"mentee_id,omitempty"`
	FullName       string `json:"full_name,omitempty"`
	Class          string `json:"class,omitempty"`
	EducationMajor string `json:"major,omitempty"`
	Email          string `json:"email,omitempty"`
	Institution    string `json:"institution,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Telegram       string `json:"telegram,omitempty"`
	Feedbacks      []feedback.CoreFeedback
}

func CoreToCreateDeleteResponse(mentee mentee.CoreMentee) CreatDeleteMenteeResponse {
	return CreatDeleteMenteeResponse{
		ID:       mentee.ID,
		FullName: mentee.FullName,
		Email:    mentee.Email,
		Phone:    mentee.Phone,
		Telegram: mentee.Telegram,
		Gender:   mentee.Gender,
	}
}

func CoreToReadUpdateResponse(mentee mentee.CoreMentee) GetUpdateMenteeResponse {
	return GetUpdateMenteeResponse{
		ID: mentee.ID,
		// StatusId:          mentee.StatusId,
		Class:             mentee.Class.ClassName,
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
		Class:         mentee.Class.ClassName,
		FullName:      mentee.FullName,
		Status:        mentee.Status.StatusName,
		EducationType: mentee.EducationType,
		Gender:        mentee.Gender,
	}
}

func CoreToGetMenteeFeedbackResponse(mentee mentee.CoreMentee) GetMenteeFeedbackResponse {
	return GetMenteeFeedbackResponse{
		ID:       mentee.ID,
		FullName: mentee.FullName,
		// ClassId:        mentee.ClassId,
		// EducationMajor: mentee.EducationMajor,
		// Institution:    mentee.Institution,
		// Phone:          mentee.Phone,
		// Telegram:       mentee.Telegram,
		// Email:          mentee.Email,
		Feedbacks: mentee.Feedbacks,
	}

}
