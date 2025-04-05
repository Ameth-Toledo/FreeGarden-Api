package use_case

import (
	"FreeGarden/src/sensor_ultrasonico/domain"
	"FreeGarden/src/sensor_ultrasonico/domain/entities"
)

type SaveUltrasonic struct {
	repo domain.UltrasonicRepository
}

func NewUltrasonicSave(repo domain.UltrasonicRepository) *SaveUltrasonic {
	return &SaveUltrasonic{repo: repo}
}

func (s *SaveUltrasonic) SaveDistance(sensor entities.UltrasonicSensor) (entities.UltrasonicSensor, error) {
	return s.repo.Save(sensor)
}
