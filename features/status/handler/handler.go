package handler

import (
	"kelompok1/immersive-dash/features/status"
)

type StatusHandler struct {
	statusService status.StatusServiceInterface
}

func New(service status.StatusServiceInterface) *StatusHandler {
	return &StatusHandler{
		statusService: service,
	}
}
