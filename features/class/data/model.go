package data

import (
	"kelompok1/immersive-dash/features/mentee/data"

	"gorm.io/gorm"
)

// struct Class gorm model
type Class struct {
	gorm.Model
	UserId      uint          `gorm:"column:user_id;not null"`
	ClassName   string        `gorm:"column:class_name;not null"`
	Mentor      string        `gorm:"column:mentor;not null"`
	StartDate   string        `gorm:"column:start_date;not null"`
	GraduteDate string        `gorm:"column:graduate_date;not null"`
	Mentees     []data.Mentee `gorm:"foreignKey:ClassID"`
}
