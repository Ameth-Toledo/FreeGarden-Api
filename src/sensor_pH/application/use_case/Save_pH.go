package use_case

import (
	"FreeGarden/src/sensor_pH/domain"
	"FreeGarden/src/sensor_pH/domain/entities"
)

type SavePH struct {
	repo domain.PHRepository
}

func NewSavePH(repo domain.PHRepository) *SavePH {
	return &SavePH{repo: repo}
}

func (s *SavePH) SaveValue(sensor entities.PhSensor) (entities.PhSensor, error) {
	return s.repo.Save(sensor)
}
