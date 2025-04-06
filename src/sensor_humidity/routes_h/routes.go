package routes_h

import (
	"FreeGarden/src/sensor_humidity/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	createController *controllers.SaveHumidityController,
	getController *controllers.GetValueHumidityController,
) {
	api := router.Group("/api/sensor-humidity")
	{
		// Ruta para guardar la humedad
		api.POST("/create", createController.SaveHumidity)
		// Ruta para obtener el valor de la humedad
		api.GET("/get/:user_id", getController.GetValue)
	}
}
