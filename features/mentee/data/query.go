package data

import (
	"errors"
	"kelompok1/immersive-dash/features/mentee"
	"log"

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
func (repo *menteeQuery) SelectAll(className, statusName, educationType string) ([]mentee.CoreMentee, error) {
	var dataMentee []Mentee
	//query join
	query := repo.db.Table("mentees").Select("mentees.*, classes.class_name As class_name, statuses.status_name AS status_name").Joins("JOIN classes ON mentees.class_id = classes.id").Joins("JOIN statuses ON mentees.status_id = statuses.id")

	if className != "" {
		query = query.Where("classes.class_name = ?", className)
	}
	if statusName != "" {
		query = query.Where("statuses.status_name = ?", statusName)
	}
	if educationType != "" {
		query = query.Where("mentees.education_type = ?", educationType)
	}
	// Lanjutkan dengan eksekusi query
	tx := query.Preload("Status").Preload("Class").Find(&dataMentee)
	if tx.Error != nil {
		return nil, errors.New(tx.Error.Error() + ", failed to get mentee")
	}

	if tx.RowsAffected == 0 && className == "" && statusName == "" && educationType == "" {
		// Jika tidak ada query dan tidak ada data, jalankan query tanpa kondisi WHERE
		tx = repo.db.Find(&dataMentee)
		if tx.Error != nil {
			return nil, errors.New(tx.Error.Error() + ", failed to get mentee")
		}
		if tx.RowsAffected == 0 {
			return nil, errors.New("Mentee not found")
		}
	}

	//Mapping Mentee to Corementee
	coreMenteeSlice := ListModelToCore(dataMentee)
	return coreMenteeSlice, nil
}

// Select implements mentee.MenteeDataInterface.
func (repo *menteeQuery) Select(menteeId uint) (mentee.CoreMentee, error) {
	log.Printf("Mencari data mentee dengan ID: %d", menteeId)
	var dataMentee Mentee
	tx := repo.db.Preload("Status").Preload("Class").Preload("Feedbacks").First(&dataMentee, menteeId)
	if tx.Error != nil {
		log.Printf("Terjadi kesalahan saat mencari mentee: %v", tx.Error)
		return mentee.CoreMentee{}, errors.New(tx.Error.Error() + ", failed to get mentee")
	}
	if tx.RowsAffected == 0 {
		log.Printf("Mentee dengan ID %d tidak ditemukan", menteeId)
		return mentee.CoreMentee{}, errors.New("Mentee not found")
	}

	//Mapping Mentee to Corementee
	coreMentee := ModelToCore(dataMentee)
	log.Printf("Mentee dengan ID %d ditemukan: %v", menteeId, coreMentee)

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
