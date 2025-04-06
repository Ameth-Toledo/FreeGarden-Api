package dependencies_h

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_humidity/application/repositories"
	"FreeGarden/src/sensor_humidity/application/use_case"
	"FreeGarden/src/sensor_humidity/infraestructure/adapters_h"
	"FreeGarden/src/sensor_humidity/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	*controllers.SaveHumidityController,
	*controllers.GetValueHumidityController,
	*repositories.ServiceNotification,
	error,
) {

	repository := adapters_h.NewMySQL(pool.DB)

	// Inicializar RabbitMQ
	rabbitMQAdapter, err := adapters_h.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
	}

	// Servicio de notificación (RabbitMQ)
	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	// Casos de uso
	createUseCase := use_case.NewSaveHumidity(repository)
	getUseCase := use_case.NewGetValueHumidity(repository)

	// Controladores
	createController := controllers.NewSaveHumidityController(createUseCase, serviceNotification)
	getController := controllers.NewGetValueHUmidityController(getUseCase)

	return createController, getController, serviceNotification, nil
}
