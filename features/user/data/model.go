package data

import (
	"kelompok1/immersive-dash/features/user"

	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	FullName string `gorm:"column:full_name;not null"`
	Email    string `gorm:"column:email;unique;not null"`
	Password string `gorm:"column:password;not null"`
	TeamId   uint   `gorm:"column:team_id;not null"`
	Role     string `gorm:"type:enum('super admin','user');default:'user';column:role;not null"`
	Status   bool   `gorm:"column:status;not null"`
}

// Mapping struct core to struct model
func CoreToModel(dataCore user.CoreUser) User {
	return User{
		FullName: dataCore.FullName,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		TeamId:   dataCore.TeamId,
		Role:     dataCore.Role,
		Status:   dataCore.Status,
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel User) user.CoreUser {
	return user.CoreUser{
		ID:        dataModel.ID,
		FullName:  dataModel.FullName,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		TeamId:    dataModel.TeamId,
		Role:      dataModel.Role,
		Status:    dataModel.Status,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
		DeletedAt: dataModel.DeletedAt.Time,
	}
}

// mapping []model to []core
func ListModelToCore(dataModel []User) []user.CoreUser {
	var result []user.CoreUser
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
