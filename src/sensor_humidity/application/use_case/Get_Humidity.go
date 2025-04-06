package use_case

import (
	"FreeGarden/src/sensor_humidity/domain"
)

type GetValueHumidity struct {
	repo domain.IHumidity
}

func NewGetValueHumidity(repo domain.IHumidity) *GetValueHumidity {
	return &GetValueHumidity{repo: repo}
}

func (s *GetValueHumidity) GetLastValueByUserID(userID int) (float64, error) {
	HumiditySensor, err := s.repo.GetLastValueByUserID(userID)
	if err != nil {
		return 0, err
	}
	return HumiditySensor.Humidity, nil
}
