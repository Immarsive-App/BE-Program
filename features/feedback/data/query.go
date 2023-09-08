package data

import (
	"errors"
	"kelompok1/immersive-dash/features/feedback"

	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

// Insert implements feedback.FeedbackDataInterface
func (repo *feedbackQuery) Insert(input feedback.CoreFeedback, userId uint) (feedback.CoreFeedback, error) {

	feedbackGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&feedbackGorm) // proses query insert
	if tx.Error != nil {
		return feedback.CoreFeedback{}, tx.Error
	}
	return feedback.CoreFeedback{}, nil
}

// Delete implements feedback.FeedbackDataInterface
func (repo *feedbackQuery) Delete(id uint) error {
	var projectGorm Feedback
	tx := repo.db.First(&projectGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	// Hapus pengguna dari database
	tx = repo.db.Delete(&projectGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found to deleted")

	}
	return nil
}

// Update implements feedback.FeedbackDataInterface
func (repo *feedbackQuery) Update(id uint, input feedback.CoreFeedback) error {
	var feedbackGorm Feedback
	tx := repo.db.First(&feedbackGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	// Perbarui properti pengguna dengan data dari input
	feedbackGorm.Note = input.Note
	feedbackGorm.UserId = input.UserId
	feedbackGorm.MenteeId = input.MenteeId
	feedbackGorm.StatusId = input.StatusId

	tx = repo.db.Save(&feedbackGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func New(db *gorm.DB) feedback.FeedbackDataInterface {
	return &feedbackQuery{
		db: db,
	}
}
