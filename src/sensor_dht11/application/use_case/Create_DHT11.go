package use_case

import (
	"FreeGarden/src/sensor_dht11/application/repositories"
	"FreeGarden/src/sensor_dht11/domain"
	"FreeGarden/src/sensor_dht11/domain/entities"
	"log"
)

type Create_DHT11 struct {
	dht11Repo           domain.IDHT11Sensor
	serviceNotification *repositories.ServiceNotification
}

func NewCreate_DHT11(dht11Repo domain.IDHT11Sensor, serviceNotification *repositories.ServiceNotification) *Create_DHT11 {
	return &Create_DHT11{
		dht11Repo:           dht11Repo,
		serviceNotification: serviceNotification,
	}
}

func (c *Create_DHT11) Execute(gyroscope entities.DHT11Sensor) (entities.DHT11Sensor, error) {
	created, err := c.dht11Repo.Save(gyroscope)

	err = c.serviceNotification.NotifyAppoinmentCreated(created)
	if err != nil {
		log.Printf("Error notificando medición creada: %v", err)
		return entities.DHT11Sensor{}, err
	}

	return created, nil
}
