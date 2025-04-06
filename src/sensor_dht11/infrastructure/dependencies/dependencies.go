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

func InitializeSensorDHTDependencies(router *gin.Engine) (*use_case.CreateDHT, *use_case.GetValueDHT, *repositories.ServiceNotification, error) {
	dbConn := core.GetDBPool()

	DHTRepo := adapters.NewMySQLDHTRepository(dbConn)

	saveDHTUseCase := use_case.NewCreateDHT(DHTRepo)
	getValueUseCase := use_case.NewGetValueDHT(DHTRepo)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error inicializando RabbitMQ: %v", err)
		return nil, nil, nil, err
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	saveDHTController := controllers.NewSaveValueController(saveDHTUseCase, serviceNotification)
	getValueController := controllers.NewGetValueDHTController(getValueUseCase)

	routes.SetupRoutes(router, saveDHTController, getValueController)

	return saveDHTUseCase, getValueUseCase, serviceNotification, nil
}
