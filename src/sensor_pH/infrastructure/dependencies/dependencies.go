package dependencies

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_pH/application/repositories"
	"FreeGarden/src/sensor_pH/application/use_case"
	"FreeGarden/src/sensor_pH/infrastructure"
	"FreeGarden/src/sensor_pH/infrastructure/adapters"
	"FreeGarden/src/sensor_pH/infrastructure/controllers"
	"FreeGarden/src/sensor_pH/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func InitializeSensorPhDependencies(
	router *gin.Engine,
) (*use_case.SavePH, *use_case.GetValuePH, *repositories.ServiceNotification, error) {
	dbConn := core.GetDBPool()

	phRepo := infrastructure.NewMySQLpHRepository(dbConn)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error inicializando RabbitMQ: %v", err)
		return nil, nil, nil, err
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	savePhUseCase := use_case.NewSavePH(phRepo)
	getValueUseCase := use_case.NewGetValuePH(phRepo)

	savePhController := controllers.NewSaveValueController(savePhUseCase)
	getValueController := controllers.NewGetValuePhController(getValueUseCase)

	routes.SetupRoutes(router, savePhController, getValueController)

	return savePhUseCase, getValueUseCase, serviceNotification, nil
}
