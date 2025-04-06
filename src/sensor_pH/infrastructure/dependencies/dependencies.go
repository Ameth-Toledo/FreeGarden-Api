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

	// Repositorio
	phRepo := infrastructure.NewMySQLpHRepository(dbConn)

	// Adaptador RabbitMQ
	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error inicializando RabbitMQ: %v", err)
		return nil, nil, nil, err
	}

	// Servicio de notificación
	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	// Casos de uso (pasando tanto el repo como el adaptador)
	savePhUseCase := use_case.NewSavePH(phRepo, serviceNotification)
	getValueUseCase := use_case.NewGetValuePH(phRepo)

	// Controladores
	savePhController := controllers.NewSaveValueController(savePhUseCase)
	getValueController := controllers.NewGetValuePhController(getValueUseCase)

	// Rutas
	routes.SetupRoutes(router, savePhController, getValueController)

	return savePhUseCase, getValueUseCase, serviceNotification, nil
}
