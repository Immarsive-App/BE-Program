package data

import (
	"gorm.io/gorm"
)

// struct Mentee gorm model
type Mentee struct {
	gorm.Model
	ClassId           uint
	StatusId          uint
	MenteeName        string
	CurrentAdress     string
	HomeAdress        string
	Email             string
	Gender            string
	Telegram          string
	Phone             string
	EmergencyName     string
	EmergencyPhone    string
	EmergencyStatus   string
	EducationType     string
	EducationMajor    string
	Institution       string
	EducationGraduate string
}
