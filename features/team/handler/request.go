package handler

import (
	"kelompok1/immersive-dash/features/team"
)

type TeamRequest struct {
	TeamName string `json:"name" form:"team_name"`
}

func RequestToCore(input TeamRequest) team.CoreTeam {
	return team.CoreTeam{
		TeamName: input.TeamName,
	}
}
