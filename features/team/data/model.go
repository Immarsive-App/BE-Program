package data

import (
	"gorm.io/gorm"
)

// struct Team gorm model
type Team struct {
	gorm.Model
	TeamName string
}
