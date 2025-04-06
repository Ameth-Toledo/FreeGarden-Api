package domain

import "FreeGarden/src/sensor_dht11/domain/entities"

type IDHT11Sensor interface {
	Save(sensor entities.DHT11Sensor) (entities.DHT11Sensor, error)
	GetLastValueByUserID(userID int) (entities.DHT11Sensor, error)
}
