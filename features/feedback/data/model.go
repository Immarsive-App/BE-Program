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

func CoreToModel(dataCore feedback.CoreFeedback) Feedback {
	return Feedback{
		Model:    gorm.Model{},
		UserId:   dataCore.UserId,
		MenteeId: dataCore.MenteeId,
		StatusId: dataCore.StatusId,
		Note:     dataCore.Note,
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel Feedback) feedback.CoreFeedback {
	return feedback.CoreFeedback{
		ID:        dataModel.ID,
		UserId:    dataModel.UserId,
		MenteeId:  dataModel.MenteeId,
		StatusId:  dataModel.StatusId,
		Note:      dataModel.Note,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func ListModelToCore(dataModel []Feedback) []feedback.CoreFeedback {
	var result []feedback.CoreFeedback
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
