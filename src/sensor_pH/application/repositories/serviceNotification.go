package repositories

import (
	"FreeGarden/src/sensor_pH/domain/entities"
	"log"
)

type ServiceNotification struct {
	notificationPort NotificationPort
}

func NewServiceNotification(notificationPort NotificationPort) *ServiceNotification {
	return &ServiceNotification{notificationPort: notificationPort}
}

// ✅ Este método es opcional
func (sn *ServiceNotification) NotifyAppointmentCreated(sensor entities.PhSensor) error {
	log.Println("Notificando la creación del sensor...")

	err := sn.notificationPort.PublishEvent("sensor_creado", sensor)
	if err != nil {
		log.Printf("Error al publicar el evento: %v", err)
		return err
	}
	return nil
}

// ✅ Este método implementa la interfaz NotificationPort
func (sn *ServiceNotification) PublishEvent(eventType string, sensor entities.PhSensor) error {
	return sn.notificationPort.PublishEvent(eventType, sensor)
}
