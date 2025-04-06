package repositories

import "FreeGarden/src/sensor_humidity/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.Humidity) error
}
