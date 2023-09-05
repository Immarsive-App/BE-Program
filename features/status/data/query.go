package data

import (
	"kelompok1/immersive-dash/features/status"

	"gorm.io/gorm"
)

type statusQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) status.StatusDataInterface {
	return &statusQuery{
		db: db,
	}
}
