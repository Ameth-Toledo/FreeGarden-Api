package dependencies

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_ultrasonico/application/repositories"
	"FreeGarden/src/sensor_ultrasonico/application/use_case"
	"FreeGarden/src/sensor_ultrasonico/infrastructure"
	"FreeGarden/src/sensor_ultrasonico/infrastructure/adapters"
	controllers "FreeGarden/src/sensor_ultrasonico/infrastructure/controllers"
	"FreeGarden/src/sensor_ultrasonico/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"log"
)

// InitializeSensorUltrasonicDependencies configura las dependencias necesarias para el sensor ultrasónico.
func InitializeSensorUltrasonicDependencies(
	router *gin.Engine,
) (*use_case.SaveUltrasonic, *use_case.GetDistanceUltrasonic, *repositories.ServiceNotification, error) {
	dbConn := core.GetDBPool()

	ultrasonicRepo := infrastructure.NewMySQLUltrasonicRepository(dbConn)

	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error inicializando RabbitMQ: %v", err)
		return nil, nil, nil, err
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	saveUltrasonicUseCase := use_case.NewUltrasonicSave(ultrasonicRepo)
	getDistanceUseCase := use_case.NewUltrasonicGetDistance(ultrasonicRepo)

	saveDistanceController := controllers.NewSaveDistanceController(saveUltrasonicUseCase)
	getDistanceController := controllers.NewGetDistanceController(getDistanceUseCase)

	routes.SetupRoutes(router, saveDistanceController, getDistanceController)

	return saveUltrasonicUseCase, getDistanceUseCase, serviceNotification, nil
}
