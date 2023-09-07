package service

import "kelompok1/immersive-dash/features/team"

type teamService struct {
	teamData team.TeamDataInterface
}

func New(repo team.TeamDataInterface) team.TeamServiceInterface {
	return &teamService{
		teamData: repo,
	}
}

// GetAll implements team.TeamServiceInterface.
func (service *teamService) GetAll() ([]team.CoreTeam, error) {
	result, err := service.teamData.SelectAll()
	return result, err
}
