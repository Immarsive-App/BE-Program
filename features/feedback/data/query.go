package data

import (
	"kelompok1/immersive-dash/features/feedback"

	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.FeedbackDataInterface {
	return &feedbackQuery{
		db: db,
	}
}
