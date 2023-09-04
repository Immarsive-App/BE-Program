package data

import (
	"kelompok1/immersive-dash/features/mentee/data"

	"gorm.io/gorm"
)

// struct Class gorm model
type Class struct {
	gorm.Model
	UserId    uint          `gorm:"column:user_id;not null"`
	ClassName string        `gorm:"column:class_name;not null"`
	Mentees   []data.Mentee `gorm:"foreignKey:ClassId"`
}
