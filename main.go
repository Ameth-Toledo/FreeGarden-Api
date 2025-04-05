package main

import (
	"FreeGarden/src/sensor_ultrasonico/infrastructure/dependencies"
	"log"
)

func main() {
	// Inicializa solo las dependencias necesarias para el sensor ultrasonico
	router, _, _ := dependencies.InitializeSensorUltrasonicDependencies()

	// Inicia el servidor en el puerto 8080 (o el puerto que prefieras)
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
