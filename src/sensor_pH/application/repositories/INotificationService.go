package repositories

import "FreeGarden/src/sensor_pH/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, sensor entities.PhSensor) error
}
