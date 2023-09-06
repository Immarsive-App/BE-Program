package handler

import (
	"kelompok1/immersive-dash/features/mentee"
)

type MenteeResponse struct {
	ID            uint   `json:"id"`
	ClassId       uint   `json:"class_id"`
	FullName      string `json:"full_name"`
	StatusId      uint   `json:"status_id"`
	EducationType string `json:"education_type"`
	Gender        string `json:"gender"`
}

func CoreToResponse(core mentee.CoreMentee) MenteeResponse {
	return MenteeResponse{
		ID:            core.ID,
		ClassId:       core.ClassId,
		FullName:      core.FullName,
		StatusId:      core.StatusId,
		EducationType: core.EducationType,
		Gender:        core.Gender,
	}
}
