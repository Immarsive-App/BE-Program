package handler

import (
	"kelompok1/immersive-dash/features/team"
)

type TeamHandler struct {
	teamService team.TeamServiceInterface
}

func New(service team.TeamServiceInterface) *TeamHandler {
	return &TeamHandler{
		teamService: service,
	}
}
