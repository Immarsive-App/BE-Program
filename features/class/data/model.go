package data

import (
	"gorm.io/gorm"
)

// struct Class gorm model
type Class struct {
	gorm.Model
	UserId      uint
	ClassName   string
	Mentor      string
	StartDate   string
	GraduteDate string
}
