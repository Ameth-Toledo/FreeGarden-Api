package use_case

import "FreeGarden/src/sensor_humidity/domain"

type GetAverageHumidity struct {
	db domain.IHumidity
}

func NewGetAverageHumidity(db domain.IHumidity) *GetAverageHumidity {
	return &GetAverageHumidity{db: db}
}

func (gt *GetAverageHumidity) Execute(userID int) (float64, error) {
	return gt.db.GetAverageHumidity(userID)
}
