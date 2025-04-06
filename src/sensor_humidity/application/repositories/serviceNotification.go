package repositories

import (
	"FreeGarden/src/sensor_humidity/domain/entities"
	"log"
)

type ServiceNotification struct {
	notificationPort NotificationPort
}

func NewServiceNotification(notificationPort NotificationPort) *ServiceNotification {
	return &ServiceNotification{notificationPort: notificationPort}
}

func (sn *ServiceNotification) NotifyAppointmentCreated(appointment entities.Humidity) error {
	log.Println("Creando dato humedad...")

	err := sn.notificationPort.PublishEvent("cita_creada", appointment)
	if err != nil {
		log.Printf("Error al publicar el evento: %v", err)
		return err
	}
	return nil
}
