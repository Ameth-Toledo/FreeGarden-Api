package use_case

import (
	"FreeGarden/src/sensor_humidity/domain"
	"FreeGarden/src/sensor_humidity/domain/entities"
)

type Get_Latest struct {
	db domain.IHumidity
}

func NewGet_Latest(db domain.IHumidity) *Get_Latest {
	return &Get_Latest{db: db}
}

func (gt *Get_Latest) Execute(userID int) (entities.Humidity, error) {
	return gt.db.GetLatestMeasurement(userID)
}
