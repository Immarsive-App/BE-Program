package data

import (
	"kelompok1/immersive-dash/features/team"

	"gorm.io/gorm"
)

type teamQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) team.TeamDataInterface {
	return &teamQuery{
		db: db,
	}
}
