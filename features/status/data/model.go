package data

import (
	"kelompok1/immersive-dash/features/feedback"
	_feedbackData "kelompok1/immersive-dash/features/feedback/data"
	"kelompok1/immersive-dash/features/status"

	"gorm.io/gorm"
)

// struct Status gorm model
type Status struct {
	gorm.Model
	StatusName string                   `gorm:"status_name;unique;not null"`
	Feedbacks  []_feedbackData.Feedback `gorm:"foreignKey:StatusId"`
}

func CoreToModel(dataCore status.CoreStatus) Status {
	return Status{
		Model:      gorm.Model{},
		StatusName: dataCore.StatusName,
		Feedbacks:  []_feedbackData.Feedback{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel Status) status.CoreStatus {
	return status.CoreStatus{
		ID:         dataModel.ID,
		StatusName: dataModel.StatusName,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
		Feedbacks:  []feedback.CoreFeedback{},
	}
}

func ListModelToCore(dataModel []Status) []status.CoreStatus {
	var result []status.CoreStatus
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
