package dependencies

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_dht11/application/repositories"
	"FreeGarden/src/sensor_dht11/application/use_case"
	"FreeGarden/src/sensor_dht11/infrastructure/adapters"
	"FreeGarden/src/sensor_dht11/infrastructure/controllers"
	"FreeGarden/src/sensor_dht11/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"log"
)

// InitializeSensorDHTDependencies configura las dependencias necesarias para el sensor DHT11.
func InitializeSensorDHTDependencies() (*gin.Engine, *use_case.CreateDHT, *use_case.GetValueDHT, *repositories.ServiceNotification, error) {
	dbConn := core.GetDBPool()

	DHTRepo := adapters.NewMySQLDHTRepository(dbConn)

	saveDHTUseCase := use_case.NewCreateDHT(DHTRepo)
	getValueUseCase := use_case.NewGetValueDHT(DHTRepo)

	// Crear el adaptador de RabbitMQ
	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error inicializando RabbitMQ: %v", err)
		return nil, nil, nil, nil, err
	}

	// Crear el servicio de notificación con RabbitMQ
	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	saveDHTController := controllers.NewSaveValueController(saveDHTUseCase)
	getValueController := controllers.NewGetValueDHTController(getValueUseCase)

	router := gin.Default()

	routes.SetupRoutes(router, saveDHTController, getValueController)

	return router, saveDHTUseCase, getValueUseCase, serviceNotification, nil
}
