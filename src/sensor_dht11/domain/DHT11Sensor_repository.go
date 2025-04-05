package domain

import "FreeGarden/src/sensor_dht11/domain/entities"

type IDHT11Sensor interface {
	Save(sensor entities.DHT11Sensor) (entities.DHT11Sensor, error)
	GetMeasurementByID(id int, userID int) (entities.DHT11Sensor, error)
	GetLatestMeasurement(userID int) (entities.DHT11Sensor, error)
	GetAllMeasurements(userID int) ([]entities.DHT11Sensor, error)
	DeleteMeasurement(id int, userID int) error
	GetAverageTemperature(userID int) (float64, error)
	GetAverageHumidity(userID int) (float64, error)
}
