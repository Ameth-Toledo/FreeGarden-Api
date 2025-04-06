package use_case

import (
	"FreeGarden/src/sensor_humidity/application/repositories"
	"FreeGarden/src/sensor_humidity/domain"
	"FreeGarden/src/sensor_humidity/domain/entities"
	"log"
)

type CreateHumidity struct {
	humidity     domain.IHumidity
	notification *repositories.ServiceNotification
}

func NewCreate_Humidity(humidity domain.IHumidity, notification *repositories.ServiceNotification) *CreateHumidity {
	return &CreateHumidity{humidity: humidity, notification: notification}
}

func (c *CreateHumidity) Execute(appointment entities.Humidity) (entities.Humidity, error) {

	created, err := c.humidity.Save(appointment)

	err = c.notification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notificando cita creada: %v", err)
		return entities.Humidity{}, err
	}

	return created, nil
}
