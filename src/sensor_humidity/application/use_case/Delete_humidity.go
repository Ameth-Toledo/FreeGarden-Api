package use_case

import "FreeGarden/src/sensor_humidity/domain"

type Delete_Humidty struct {
	db domain.IHumidity
}

func NewDelete_Humidty(db domain.IHumidity) *Delete_Humidty {
	return &Delete_Humidty{db: db}
}

func (dt *Delete_Humidty) Execute(id int, userID int) error {
	return dt.db.DeleteMeasurement(id, userID)
}
