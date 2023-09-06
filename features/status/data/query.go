package data

import (
	"kelompok1/immersive-dash/features/status"

	"gorm.io/gorm"
)

type statusQuery struct {
	db *gorm.DB
}

// SelectAll implements status.StatusDataInterface
func (repo *statusQuery) SelectAll() ([]status.CoreStatus, error) {
	var statusData []Status
	tx := repo.db.Find(&statusData) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping dari struct gorm model ke struct core (entity)
	var statusCore = ListModelToCore(statusData)

	return statusCore, nil
}

func New(db *gorm.DB) status.StatusDataInterface {
	return &statusQuery{
		db: db,
	}
}
