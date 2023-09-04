package data

import (
	"gorm.io/gorm"
)

// struct Feedback gorm model
type Feedback struct {
	gorm.Model
	UserId   uint
	MenteeId uint
	StatusId uint
	Note     string
}
