package use_case

import (
	"FreeGarden/src/sensor_dht11/domain"
	"FreeGarden/src/sensor_dht11/domain/entities"
)

type CreateDHT struct {
	repo domain.IDHT11Sensor
}

func NewCreateDHT(repo domain.IDHT11Sensor) *CreateDHT {
	return &CreateDHT{repo: repo}
}

func (s *CreateDHT) SaveValue(sensor entities.DHT11Sensor) (entities.DHT11Sensor, error) {
	return s.repo.Save(sensor)
}
