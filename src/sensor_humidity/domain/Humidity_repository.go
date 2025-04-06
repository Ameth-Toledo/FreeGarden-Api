package domain

import "FreeGarden/src/sensor_humidity/domain/entities"

type IHumidity interface {
	Save(sensor entities.Humidity) (entities.Humidity, error)
	GetMeasurementByID(id int, userID int) (entities.Humidity, error)
	GetLatestMeasurement(userID int) (entities.Humidity, error)
	GetAllMeasurements(userID int) ([]entities.Humidity, error)
	DeleteMeasurement(id int, userID int) error
	GetAverageHumidity(userID int) (float64, error)
}
