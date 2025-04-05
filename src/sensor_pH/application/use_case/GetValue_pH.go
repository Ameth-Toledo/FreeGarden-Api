package use_case

import (
	"FreeGarden/src/sensor_pH/domain"
)

type GetValuePH struct {
	repo domain.PHRepository
}

func NewGetValuePH(repo domain.PHRepository) *GetValuePH {
	return &GetValuePH{repo: repo}
}

func (s *GetValuePH) GetLastValueByUserID(userID int) (float64, error) {
	phSensor, err := s.repo.GetLastValueByUserID(userID)
	if err != nil {
		return 0, err
	}
	return phSensor.PhValue, nil
}
