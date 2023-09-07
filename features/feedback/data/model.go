package data

import (
	"kelompok1/immersive-dash/features/feedback"

	"gorm.io/gorm"
)

// struct Feedback gorm model
type Feedback struct {
	gorm.Model
	UserId   uint   `gorm:"column:user_id;not null"`
	MenteeId uint   `gorm:"column:mentee_id;not null"`
	StatusId uint   `gorm:"column:status_id;not null"`
	Note     string `gorm:"column:note;not null"`
}

type GetMenteeFeedback struct {
	ID       uint   `gorm:"column:feedback_id;not null"`
	UserId   uint   `gorm:"column:user_id;not null"`
	MenteeId uint   `gorm:"column:mentee_id;not null"`
	StatusId uint   `gorm:"column:status_id;not null"`
	Note     string `gorm:"column:note;not null"`
}

func CoreToModel(core feedback.CoreFeedback) Feedback {
	return Feedback{
		UserId:   core.UserId,
		MenteeId: core.MenteeId,
		StatusId: core.StatusId,
		Note:     core.Note,
	}
}

func ModelToCore(model Feedback) feedback.CoreFeedback {
	return feedback.CoreFeedback{
		ID:        model.ID,
		UserId:    model.UserId,
		MenteeId:  model.MenteeId,
		StatusId:  model.StatusId,
		Note:      model.Note,
		CreatedAt: model.CreatedAt,
	}
}
