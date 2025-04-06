package use_case

import (
	"FreeGarden/src/sensor_pH/application/repositories"
	"FreeGarden/src/sensor_pH/domain"
	"FreeGarden/src/sensor_pH/domain/entities"
)

type SavePH struct {
	repo          domain.PHRepository
	rabbitAdapter repositories.NotificationPort
}

func NewSavePH(repo domain.PHRepository, rabbitAdapter repositories.NotificationPort) *SavePH {
	return &SavePH{
		repo:          repo,
		rabbitAdapter: rabbitAdapter,
	}
}

func (s *SavePH) SaveValue(sensor entities.PhSensor) (entities.PhSensor, error) {
	savedSensor, err := s.repo.Save(sensor)
	if err != nil {
		return savedSensor, err
	}

	err = s.rabbitAdapter.PublishEvent("SensorData", savedSensor)
	if err != nil {
		return savedSensor, err
	}

	return savedSensor, nil
}
