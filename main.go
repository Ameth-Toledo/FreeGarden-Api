package main

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_humidity/infraestructure/dependencies_h"
	"FreeGarden/src/sensor_humidity/routes_h"
	"FreeGarden/src/sensor_ultrasonico/infrastructure/dependencies"
	"log"
)

func main() {

	pool := core.GetDBPool()

	createController,
		getHumidityByIDController,
		getAllController,
		deleteController,
		getAverageHumidityController,
		getLatestMeasurementController,
		_,
		err := dependencies_h.Init(pool)

	if err != nil {
		log.Fatalf("Error al inicializar dependencias del sensor de humedad: %v", err)
	}

	router, _, _ := dependencies.InitializeSensorUltrasonicDependencies()

	routes_h.RegisterHumidityRoutes(
		router,
		createController,
		getHumidityByIDController,
		getAllController,
		deleteController,
		getAverageHumidityController,
		getLatestMeasurementController,
	)
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
