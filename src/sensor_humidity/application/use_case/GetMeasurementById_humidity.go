package use_case

import "FreeGarden/src/sensor_humidity/domain"

type GetMeasurementByID struct {
	db domain.IHumidity
}

func NewGetMeasurementByID(db domain.IHumidity) *GetMeasurementByID {
	return &GetMeasurementByID{db: db}
}

func (u *GetMeasurementByID) Execute(id int, userID int) (interface{}, error) {
	return u.db.GetMeasurementByID(id, userID)
}
