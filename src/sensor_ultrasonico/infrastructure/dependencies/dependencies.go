package dependencies

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_ultrasonico/application/use_case"
	"FreeGarden/src/sensor_ultrasonico/infrastructure"
	"FreeGarden/src/sensor_ultrasonico/infrastructure/controllers"
	"FreeGarden/src/sensor_ultrasonico/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

// InitializeSensorUltrasonicDependencies configura las dependencias necesarias para el sensor ultrasonico.
func InitializeSensorUltrasonicDependencies() (*gin.Engine, *use_case.SaveUltrasonic, *use_case.GetDistanceUltrasonic) {
	// Obtener la conexión a la base de datos
	dbConn := core.GetDBPool()

	// Crear el repositorio de ultrasonic con la conexión MySQL
	ultrasonicRepo := infrastructure.NewMySQLUltrasonicRepository(dbConn)

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

	// Devolver solo las dependencias necesarias
	return router, saveUltrasonicUseCase, getDistanceUseCase
}
