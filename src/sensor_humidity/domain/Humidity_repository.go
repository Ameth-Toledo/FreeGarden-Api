package domain

import "FreeGarden/src/sensor_humidity/domain/entities"

type IHumidity interface {
	Save(sensor entities.Humidity) (entities.Humidity, error)
	GetLastValueByUserID(userID int) (entities.Humidity, error)
}
