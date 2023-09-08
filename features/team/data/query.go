package data

import (
	"fmt"
	"kelompok1/immersive-dash/features/team"

	"gorm.io/gorm"
)

type teamQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) team.TeamDataInterface {
	return &teamQuery{
		db: db,
	}
}

// SelectAll implements team.TeamDataInterface.
func (repo *teamQuery) SelectAll() ([]team.CoreTeam, error) {
	var teamsData []Team
	tx := repo.db.Find(&teamsData) // select * from teams;
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println("teams:", teamsData)
	//mapping dari struct gorm model ke struct CoreUser (entity)
	var teamsCoreTeam []team.CoreTeam
	for _, value := range teamsData {
		var team = team.CoreTeam{
			ID:        value.ID,
			TeamName:  value.TeamName,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		teamsCoreTeam = append(teamsCoreTeam, team)
	}
	return teamsCoreTeam, nil
}
