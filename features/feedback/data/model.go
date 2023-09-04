package data

import (
	"gorm.io/gorm"
)

// struct Feedback gorm model
type Feedback struct {
	gorm.Model
	UserId   uint   `gorm:"column:user_id;not null"`
	MenteeId uint   `gorm:"column:mentee_id;not null"`
	StatusId uint   `gorm:"column:status_id;not null"`
	Note     string `gorm:"column:note;not null"`
}
