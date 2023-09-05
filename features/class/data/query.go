package data

import (
	"errors"
	"kelompok1/immersive-dash/features/class"

	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// Delete implements class.ClassDataInterface
func (repo *classQuery) Delete(id uint) error {
	var classGorm Class
	tx := repo.db.First(&classGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	// Hapus pengguna dari database
	tx = repo.db.Delete(&classGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found to deleted")

	}
	return nil
}

// Update implements class.ClassDataInterface
func (repo *classQuery) Update(id uint, input class.CoreClass) error {
	var classGorm Class
	tx := repo.db.First(&classGorm, id)
	if tx.Error != nil {
		return tx.Error
	}
	// Perbarui properti pengguna dengan data dari input
	classGorm.ClassName = input.ClassName
	classGorm.UserId = input.UserId
	tx = repo.db.Save(&classGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

// SelectById implements class.ClassDataInterface
func (repo *classQuery) SelectById(id uint) (class.CoreClass, error) {
	var result Class
	tx := repo.db.First(&result, id)
	if tx.Error != nil {
		return class.CoreClass{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return class.CoreClass{}, errors.New("data not found")
	}

	resultCore := ModelToCore(result)
	return resultCore, nil
}

// Insert implements class.ClassDataInterface
func (repo *classQuery) Insert(input class.CoreClass) error {
	classGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&classGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAll implements class.ClassDataInterface
func (repo *classQuery) SelectAll() ([]class.CoreClass, error) {
	var classData []Class
	tx := repo.db.Find(&classData) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping dari struct gorm model ke struct core (entity)
	var classCore = ListModelToCore(classData)

	return classCore, nil
}

func New(db *gorm.DB) class.ClassDataInterface {
	return &classQuery{
		db: db,
	}
}
