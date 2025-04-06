package use_case

import (
	"FreeGarden/src/sensor_humidity/domain"
	"FreeGarden/src/sensor_humidity/domain/entities"
)

type Get_All_Humidity struct {
	db domain.IHumidity
}

func NewGet_All(db domain.IHumidity) *Get_All_Humidity {
	return &Get_All_Humidity{db: db}
}

func (gt *Get_All_Humidity) Execute(userID int) ([]entities.Humidity, error) {
	return gt.db.GetAllMeasurements(userID)
}
