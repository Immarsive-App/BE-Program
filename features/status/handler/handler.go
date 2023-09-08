package handler

import (
	"kelompok1/immersive-dash/features/status"
	"kelompok1/immersive-dash/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	statusService status.StatusServiceInterface
}

func New(service status.StatusServiceInterface) *StatusHandler {
	return &StatusHandler{
		statusService: service,
	}
}

func (handler *StatusHandler) GetAllStatus(c echo.Context) error {
	result, err := handler.statusService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct core to struct response
	var statusResponse []StatusResponse
	for _, value := range result {
		statusResponse = append(statusResponse, StatusResponse{
			ID:         value.ID,
			StatusName: value.StatusName,
			CreatedAt:  value.CreatedAt,
			UpdatedAt:  value.UpdatedAt,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", statusResponse))

}
