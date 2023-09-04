package database

import (
	_classData "kelompok1/immersive-dash/features/class/data"
	_feedbackData "kelompok1/immersive-dash/features/feedback/data"
	_menteeData "kelompok1/immersive-dash/features/mentee/data"
	_statusData "kelompok1/immersive-dash/features/status/data"
	_teamData "kelompok1/immersive-dash/features/team/data"
	_userData "kelompok1/immersive-dash/features/user/data"

	"gorm.io/gorm"
)

// db migration
func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
	db.AutoMigrate(&_teamData.Team{})
	db.AutoMigrate(&_classData.Class{})
	db.AutoMigrate(&_menteeData.Mentee{})
	db.AutoMigrate(&_statusData.Status{})
	db.AutoMigrate(&_feedbackData.Feedback{})
	/*
		TODO 2:
		migrate struct item
	*/
}
