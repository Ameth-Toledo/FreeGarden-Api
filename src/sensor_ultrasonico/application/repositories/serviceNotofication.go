package repositories

import (
	"FreeGarden/src/sensor_ultrasonico/domain/entities"
	"log"
)

type ServiceNotification struct {
	notificationPort NotificationPort
}

func NewServiceNotification(notificationPort NotificationPort) *ServiceNotification {
	return &ServiceNotification{notificationPort: notificationPort}
}

func (sn *ServiceNotification) NotifyAppointmentCreated(sensor entities.UltrasonicSensor) error {
	log.Println("Notificando la creación del sensor...")

	err := sn.notificationPort.PublishEvent("sensor_creado", sensor)
	if err != nil {
		log.Printf("Error al publicar el evento: %v", err)
		return err
	}
	return nil
}
