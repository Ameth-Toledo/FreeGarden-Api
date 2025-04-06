package repositories

import "FreeGarden/src/sensor_ultrasonico/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, sensor entities.UltrasonicSensor) error
}
