package main

import (
	"FreeGarden/src/core"

	dependencies_dht "FreeGarden/src/sensor_dht11/infrastructure/dependencies"
	"FreeGarden/src/sensor_humidity/infraestructure/dependencies_h"
	dependencies_ph "FreeGarden/src/sensor_pH/infrastructure/dependencies"
	dependencies_ultra "FreeGarden/src/sensor_ultrasonico/infrastructure/dependencies"

	"FreeGarden/src/sensor_humidity/routes_h"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	dbPool := core.GetDBPool()
	router := gin.Default()

	saveHumidityCtrl, getHumidityCtrl, _, err := dependencies_h.Init(dbPool)
	if err != nil {
		log.Fatalf("Error sensor humedad: %v", err)
	}
	routes_h.SetupRoutes(router, saveHumidityCtrl, getHumidityCtrl)

	_, _, _, err = dependencies_ultra.InitializeSensorUltrasonicDependencies(router)
	if err != nil {
		log.Fatalf("Error sensor ultrasonico: %v", err)
	}

	_, _, _, err = dependencies_ph.InitializeSensorPhDependencies(router)
	if err != nil {
		log.Fatalf("Error sensor pH: %v", err)
	}

	_, _, _, err = dependencies_dht.InitializeSensorDHTDependencies(router)
	if err != nil {
		log.Fatalf("Error sensor DHT11: %v", err)
	}

	if err := router.Run(":8088"); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
