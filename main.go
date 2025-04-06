package main

import (
	"FreeGarden/src/core"

	// Humedad
	"FreeGarden/src/sensor_humidity/infraestructure/dependencies_h"
	"FreeGarden/src/sensor_humidity/routes_h"

	// pH
	dependencies_ph "FreeGarden/src/sensor_pH/infrastructure/dependencies"

	// Ultrasonido
	dependencies_ultra "FreeGarden/src/sensor_ultrasonico/infrastructure/dependencies"

	// DHT11
	dependencies_dht "FreeGarden/src/sensor_dht11/infrastructure/dependencies"

	"log"
)

func main() {
	pool := core.GetDBPool()

	// Inicializar dependencias del sensor de humedad
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

	// Inicializar dependencias del sensor ultrasónico
	router, _, _, _, err := dependencies_ultra.InitializeSensorUltrasonicDependencies()
	if err != nil {
		log.Fatalf("Error al inicializar dependencias del sensor ultrasónico: %v", err)
	}

	// Inicializar dependencias del sensor de pH
	router, _, _ = dependencies_ph.InitializeSensorPhDependencies()

	// Inicializar dependencias del sensor DHT11
	router, _, _, _, err = dependencies_dht.InitializeSensorDHTDependencies()
	if err != nil {
		log.Fatalf("Error al inicializar dependencias del sensor DHT11: %v", err)
	}

	// Registrar rutas del sensor de humedad
	routes_h.RegisterHumidityRoutes(
		router,
		createController,
		getHumidityByIDController,
		getAllController,
		deleteController,
		getAverageHumidityController,
		getLatestMeasurementController,
	)

	// Iniciar servidor en el puerto 8080
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
