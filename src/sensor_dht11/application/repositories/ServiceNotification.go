package repositories

import (
	"FreeGarden/src/sensor_dht11/domain/entities"
	"log"
)

type ServiceNotification struct {
	notificationPort NotificationPort
}

func NewServiceNotification(notificationPort NotificationPort) *ServiceNotification {
	return &ServiceNotification{notificationPort: notificationPort}
}

func (notification *ServiceNotification) NotifyAppoinmentCreated(sensor entities.DHT11Sensor) error {
	log.Println("Notificando")

	err := notification.notificationPort.PublishEvent("sensor creada", sensor)
	if err != nil {
		log.Printf("Error al publicar el evento: %v", err)
		return err
	}
	return nil
}
