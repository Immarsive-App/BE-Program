package data

import (
	"gorm.io/gorm"
)

// struct Status gorm model
type Status struct {
	gorm.Model
	StatusName string `gorm:"status_name;unique;not null"`
}
