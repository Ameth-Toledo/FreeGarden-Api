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

func InitializeSensorPhDependencies() (
	*gin.Engine,
	*use_case.SavePH,
	*use_case.GetValuePH,
	*repositories.ServiceNotification,
	error,
) {
	// Obtener conexión a la base de datos
	dbConn := core.GetDBPool()

	// Crear el repositorio MySQL para pH
	phRepo := infrastructure.NewMySQLpHRepository(dbConn)

	// Crear el adaptador de RabbitMQ
	rabbitMQAdapter, err := adapters.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error inicializando RabbitMQ: %v", err)
		return nil, nil, nil, nil, err
	}

	// Crear el servicio de notificación con RabbitMQ
	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	// Crear casos de uso
	savePhUseCase := use_case.NewSavePH(phRepo)
	getValueUseCase := use_case.NewGetValuePH(phRepo)

	// Crear controladores
	savePhController := controllers.NewSaveValueController(savePhUseCase)
	getValueController := controllers.NewGetValuePhController(getValueUseCase)

	// Configurar rutas
	router := gin.Default()
	routes.SetupRoutes(router, savePhController, getValueController)

	// Retornar todo lo necesario
	return router, savePhUseCase, getValueUseCase, serviceNotification, nil
}
