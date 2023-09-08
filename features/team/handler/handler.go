package handler

import (
	"kelompok1/immersive-dash/features/team"
	"kelompok1/immersive-dash/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TeamHandler struct {
	teamService team.TeamServiceInterface
}

func New(service team.TeamServiceInterface) *TeamHandler {
	return &TeamHandler{
		teamService: service,
	}
}

func (handler *TeamHandler) GetAllTeam(c echo.Context) error {
	result, err := handler.teamService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct core to struct response
	var teamResponse []TeamResponse
	for _, value := range result {
		teamResponse = append(teamResponse, TeamResponse{
			ID:        value.ID,
			TeamName:  value.TeamName,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", teamResponse))

}
