package data

import (
	"gorm.io/gorm"
)

// struct Mentee gorm model
type Mentee struct {
	gorm.Model
	ClassId           uint   `gorm:"column:class_id;not null"`
	StatusId          uint   `gorm:"column:status_id;not null"`
	FullName          string `gorm:"column:full_name;not null"`
	CurrentAdress     string `gorm:"column:current_address; not null"`
	HomeAdress        string `gorm:"column:home_address; not null"`
	Email             string `gorm:"column:email;unique; not null"`
	Gender            string `gorm:"type:enum('male','female');default:'male';column:gender;not null"`
	Telegram          string `gorm:"column:telegram;unique; not null"`
	Phone             string `gorm:"column:phone; not null"`
	EmergencyName     string `gorm:"column:emergency_name; not null"`
	EmergencyPhone    string `gorm:"column:emergency_phone; not null"`
	EmergencyStatus   string `gorm:"type:enum('parent','grandparents', 'siblings');default:'parent';column:emergency_status;not null"`
	EducationType     string `gorm:"type:enum('informatics','non-informatics');default:'informatics';column:education_type;not null"`
	EducationMajor    string `gorm:"column:major; not null"`
	Institution       string `gorm:"column:institution; not null"`
	EducationGraduate string `gorm:"column:education_graduate; not null"`
}
