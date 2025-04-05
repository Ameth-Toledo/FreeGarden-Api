package domain

import "FreeGarden/src/sensor_dth11/domain/entities"

type IDTH11Sensor interface {
	Save(sensor entities.DTH11Sensor) (entities.DTH11Sensor, error)
	GetMeasurementByID(id int, userID int) (entities.DTH11Sensor, error)
	GetLatestMeasurement(userID int) (entities.DTH11Sensor, error)
	GetAllMeasurements(userID int) ([]entities.DTH11Sensor, error)
	DeleteMeasurement(id int, userID int) error
	GetAverageTemperature(userID int) (float64, error)
	GetAverageHumidity(userID int) (float64, error)
}
