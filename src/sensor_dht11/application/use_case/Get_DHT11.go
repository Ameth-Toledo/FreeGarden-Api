package use_case

import (
	"FreeGarden/src/sensor_dht11/domain"
)

type GetValueDHT struct {
	repo domain.IDHT11Sensor
}

func NewGetValueDHT(repo domain.IDHT11Sensor) *GetValueDHT {
	return &GetValueDHT{repo: repo}
}

func (s *GetValueDHT) GetLastValueByUserID(userID int) (float64, error) {
	dhtSensor, err := s.repo.GetLastValueByUserID(userID)
	if err != nil {
		return 0, err
	}
	return dhtSensor.Temperature, nil
}
