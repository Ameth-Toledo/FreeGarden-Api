package dependencies

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_ultrasonico/application/repositories"
	"FreeGarden/src/sensor_ultrasonico/application/use_case"
	"FreeGarden/src/sensor_ultrasonico/infrastructure"
	"FreeGarden/src/sensor_ultrasonico/infrastructure/adapters"
	"FreeGarden/src/sensor_ultrasonico/infrastructure/controllers"
	"FreeGarden/src/sensor_ultrasonico/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"log"
)

// InitializeSensorUltrasonicDependencies configura las dependencias necesarias para el sensor ultrasónico.
func InitializeSensorUltrasonicDependencies() (*gin.Engine, *use_case.SaveUltrasonic, *use_case.GetDistanceUltrasonic, *repositories.ServiceNotification, error) {
	// Obtener la conexión a la base de datos
	dbConn := core.GetDBPool()

	// Crear el repositorio de ultrasonido con la conexión MySQL
	ultrasonicRepo := infrastructure.NewMySQLUltrasonicRepository(dbConn)

	// Crear el adaptador de RabbitMQ
	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error inicializando RabbitMQ: %v", err)
		return nil, nil, nil, nil, err
	}

	// Crear el servicio de notificación con RabbitMQ
	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	// Crear los casos de uso
	saveUltrasonicUseCase := use_case.NewUltrasonicSave(ultrasonicRepo)
	getDistanceUseCase := use_case.NewUltrasonicGetDistance(ultrasonicRepo)

	// Crear los controladores
	saveDistanceController := controller.NewSaveDistanceController(saveUltrasonicUseCase)
	getDistanceController := controller.NewGetDistanceController(getDistanceUseCase)

	// Crear el router de Gin
	router := gin.Default()

	// Configurar las rutas a través de la función SetupRoutes en routes
	routes.SetupRoutes(router, saveDistanceController, getDistanceController)

	// Devolver todas las dependencias necesarias
	return router, saveUltrasonicUseCase, getDistanceUseCase, serviceNotification, nil
}
