package data

import (
	"kelompok1/immersive-dash/features/class"

	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.ClassDataInterface {
	return &classQuery{
		db: db,
	}
}
