package data

import (
	"kelompok1/immersive-dash/features/team"
	"kelompok1/immersive-dash/features/user"
	"kelompok1/immersive-dash/features/user/data"

	"gorm.io/gorm"
)

// struct Team gorm model
type Team struct {
	gorm.Model
	TeamName string      `gorm:"team_name;unique;not null"`
	Users    []data.User `gorm:"foreignKey:TeamId"`
}

func CoreToModel(dataCore team.CoreTeam) Team {
	return Team{
		Model:    gorm.Model{},
		TeamName: dataCore.TeamName,
		Users:    []data.User{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel Team) team.CoreTeam {
	return team.CoreTeam{
		ID:        dataModel.ID,
		TeamName:  dataModel.TeamName,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
		Users:     []user.CoreUser{},
	}
}

func ListModelToCore(dataModel []Team) []team.CoreTeam {
	var result []team.CoreTeam
	for _, v := range dataModel {
		result = append(result, ModelToCore(v))
	}
	return result
}
