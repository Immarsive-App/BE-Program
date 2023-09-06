package handler

import (
	"kelompok1/immersive-dash/features/status"
)

type StatusRequest struct {
	StatusName string `json:"name" form:"status_name"`
}

func RequestToCore(input StatusRequest) status.CoreStatus {
	return status.CoreStatus{
		StatusName: input.StatusName,
	}
}
