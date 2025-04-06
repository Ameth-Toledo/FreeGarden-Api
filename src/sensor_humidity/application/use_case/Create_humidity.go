package use_case

import (
	"FreeGarden/src/sensor_humidity/domain"
	"FreeGarden/src/sensor_humidity/domain/entities"
)

type SaveHumidity struct {
	repo domain.IHumidity
}

func NewSaveHumidity(repo domain.IHumidity) *SaveHumidity {
	return &SaveHumidity{repo: repo}
}

func (s *SaveHumidity) SaveValue(sensor entities.Humidity) (entities.Humidity, error) {
	return s.repo.Save(sensor)
}
