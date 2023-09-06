package data

import (
	"errors"
	"kelompok1/immersive-dash/features/mentee"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.MenteeDataInterface {
	return &menteeQuery{
		db: db,
	}
}

// Insert implements mentee.MenteeDataInterface.
func (repo *menteeQuery) Insert(input mentee.CoreMentee) (uint, error) {
	menteeModel := CoreToModel(input)

	//simpan ke db
	tx := repo.db.Create(&menteeModel)
	if tx.Error != nil {
		return 0, errors.New(tx.Error.Error() + ", failed to create mentee")
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("data not found")
	}
	return menteeModel.ID, nil

}

// SelectAll implements mentee.MenteeDataInterface.
func (repo *menteeQuery) SelectAll() ([]mentee.CoreMentee, error) {
	var dataMentee []Mentee
	tx := repo.db.Find(&dataMentee)
	if tx.Error != nil {
		return nil, errors.New(tx.Error.Error() + ", failed to get mentee")
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("Mentee not found")
	}
	//Mapping Mentee to Corementee
	coreMenteeSlice := ListModelToCore(dataMentee)
	return coreMenteeSlice, nil
}

// Select implements mentee.MenteeDataInterface.
func (repo *menteeQuery) Select(menteeId uint) (mentee.CoreMentee, error) {
	var dataMentee Mentee
	tx := repo.db.Preload("Feedbacks").First(&dataMentee, menteeId)
	if tx.Error != nil {
		return mentee.CoreMentee{}, errors.New(tx.Error.Error() + ", failed to get mentee")
	}
	if tx.RowsAffected == 0 {
		return mentee.CoreMentee{}, errors.New("Mentee not found")
	}
	//Mapping Mentee to Corementee
	coreMentee := ModelToCore(dataMentee)
	return coreMentee, nil
}

// Update implements mentee.MenteeDataInterface.
func (repo *menteeQuery) Update(menteeId uint, updatedMentee mentee.CoreMentee) error {
	var mentee Mentee
	tx := repo.db.First(&mentee, menteeId)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + ", failed to get mentee")
	}
	if tx.RowsAffected == 0 {
		return errors.New("Mentee not found")
	}
	input := CoreToModel(updatedMentee)
	tx = repo.db.Model(&mentee).Updates(input)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + ", failed to updated mentee")
	}

	return nil
}

// Delete implements mentee.MenteeDataInterface.
func (repo *menteeQuery) Delete(menteeId uint) error {
	tx := repo.db.Delete(&Mentee{}, menteeId)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + ", failed to deactive mentee")
	}
	if tx.RowsAffected == 0 {
		return errors.New("Mentee not found")
	}
	return nil
}
