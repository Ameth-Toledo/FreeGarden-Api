package repositories

import "FreeGarden/src/sensor_dht11/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.DHT11Sensor) error
}
