package routes

import (
	"FreeGarden/src/sensor_ultrasonico/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura las rutas de la API
func SetupRoutes(router *gin.Engine, saveDistanceController *controller.SaveDistanceController, getDistanceController *controller.GetDistanceController) {
	api := router.Group("/api/sensor-ultrasonic")
	{
		// Ruta para guardar el sensor ultrasónico
		api.POST("/create", saveDistanceController.SaveDistance)

		// Ruta para obtener la distancia del sensor
		api.GET("/get/:user_id", getDistanceController.GetDistance)
	}
}
