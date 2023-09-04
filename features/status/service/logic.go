package service

import "kelompok1/immersive-dash/features/status"

type statusService struct {
	statusData status.StatusDataInterface
}

func New(repo status.StatusDataInterface) status.StatusServiceInterface {
	return &statusService{
		statusData: repo,
	}
}
