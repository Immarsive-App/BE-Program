package service

import "kelompok1/immersive-dash/features/status"

type statusService struct {
	statusData status.StatusDataInterface
}

// GetAll implements status.StatusServiceInterface
func (service *statusService) GetAll() ([]status.CoreStatus, error) {
	result, err := service.statusData.SelectAll()
	return result, err
}

func New(repo status.StatusDataInterface) status.StatusServiceInterface {
	return &statusService{
		statusData: repo,
	}
}
