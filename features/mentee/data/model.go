package data

import (
	"kelompok1/immersive-dash/features/feedback"
	"kelompok1/immersive-dash/features/feedback/data"
	"kelompok1/immersive-dash/features/mentee"

	"gorm.io/gorm"
)

// struct Mentee gorm model
type Mentee struct {
	gorm.Model
	ClassId           uint            `gorm:"column:class_id;not null"`
	StatusId          uint            `gorm:"column:status_id;not null"`
	FullName          string          `gorm:"column:full_name;not null"`
	CurrentAddress    string          `gorm:"column:current_address; not null"`
	HomeAddress       string          `gorm:"column:home_address; not null"`
	Email             string          `gorm:"column:email;unique; not null"`
	Gender            string          `gorm:"type:enum('male','female');default:'male';column:gender;not null"`
	Telegram          string          `gorm:"column:telegram;unique; not null"`
	Phone             string          `gorm:"column:phone; not null"`
	EmergencyName     string          `gorm:"column:emergency_name; not null"`
	EmergencyPhone    string          `gorm:"column:emergency_phone; not null"`
	EmergencyStatus   string          `gorm:"type:enum('parent','grandparents', 'siblings');default:'parent';column:emergency_status;not null"`
	EducationType     string          `gorm:"type:enum('informatics','non-informatics');default:'informatics';column:education_type;not null"`
	EducationMajor    string          `gorm:"column:major; not null"`
	Institution       string          `gorm:"column:institution; not null"`
	EducationGraduate string          `gorm:"column:education_graduate; not null"`
	Feedbacks         []data.Feedback `gorm:"foreignKey:MenteeId"`
}

func CoreToModel(core mentee.CoreMentee) Mentee {
	// Convert CoreFeedback to feedback models
	var feedbackModel []data.Feedback
	for _, feedback := range core.Feedbacks {
		feedbackModel = append(feedbackModel, data.CoreToModel(feedback))
	}
	return Mentee{
		ClassId:           core.ClassId,
		StatusId:          core.StatusId,
		FullName:          core.FullName,
		CurrentAddress:    core.CurrentAddress,
		HomeAddress:       core.HomeAddress,
		Gender:            core.Gender,
		Telegram:          core.Telegram,
		Phone:             core.Phone,
		EmergencyName:     core.EmergencyName,
		EmergencyPhone:    core.EmergencyPhone,
		EmergencyStatus:   core.EmergencyStatus,
		EducationType:     core.EducationType,
		EducationMajor:    core.EducationMajor,
		Institution:       core.Institution,
		EducationGraduate: core.EducationGraduate,
		Feedbacks:         feedbackModel,
	}
}

func ModelToCore(model Mentee) mentee.CoreMentee {
	// Convert Feedback model to core feedback
	var coreFeedbacks []feedback.CoreFeedback
	for _, feedback := range model.Feedbacks {
		coreFeedbacks = append(coreFeedbacks, data.ModelToCore(feedback))
	}
	return mentee.CoreMentee{
		ID:                model.ID,
		ClassId:           model.ClassId,
		StatusId:          model.StatusId,
		FullName:          model.FullName,
		CurrentAddress:    model.CurrentAddress,
		HomeAddress:       model.HomeAddress,
		Gender:            model.Gender,
		Telegram:          model.Telegram,
		Phone:             model.Phone,
		EmergencyName:     model.EmergencyName,
		EmergencyPhone:    model.EmergencyPhone,
		EmergencyStatus:   model.EmergencyStatus,
		EducationType:     model.EducationType,
		EducationMajor:    model.EducationMajor,
		Institution:       model.Institution,
		EducationGraduate: model.EducationGraduate,
		CreatedAt:         model.CreatedAt,
		UpdatedAt:         model.UpdatedAt,
		DeletedAt:         model.DeletedAt.Time,
		Feedbacks:         coreFeedbacks,
	}
}

func ListModelToCore(models []Mentee) []mentee.CoreMentee {
	// Convert Feedback model to core feedback
	var coreMentees []mentee.CoreMentee
	for _, model := range models {
		var coreFeedbacks []feedback.CoreFeedback
		for _, feedback := range model.Feedbacks {
			coreFeedbacks = append(coreFeedbacks, data.ModelToCore(feedback))
		}
		coreMentees = append(coreMentees, mentee.CoreMentee{
			ID:                model.ID,
			ClassId:           model.ClassId,
			StatusId:          model.StatusId,
			FullName:          model.FullName,
			CurrentAddress:    model.CurrentAddress,
			HomeAddress:       model.HomeAddress,
			Gender:            model.Gender,
			Telegram:          model.Telegram,
			Phone:             model.Phone,
			EmergencyName:     model.EmergencyName,
			EmergencyPhone:    model.EmergencyPhone,
			EmergencyStatus:   model.EmergencyStatus,
			EducationType:     model.EducationType,
			EducationMajor:    model.EducationMajor,
			Institution:       model.Institution,
			EducationGraduate: model.EducationGraduate,
			CreatedAt:         model.CreatedAt,
			UpdatedAt:         model.UpdatedAt,
			DeletedAt:         model.DeletedAt.Time,
			Feedbacks:         coreFeedbacks,
		})
	}
	return coreMentees
}
