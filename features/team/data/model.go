package data

import (
	"kelompok1/immersive-dash/features/user/data"

	"gorm.io/gorm"
)

// struct Team gorm model
type Team struct {
	gorm.Model
	TeamName string      `gorm:"team_name;unique;not null"`
	Users    []data.User `gorm:"foreignKey:TeamId"`
}
