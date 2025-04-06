package routes

import (
	"FreeGarden/src/sensor_dht11/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, valueController *controllers.SaveValueController, getValueController *controllers.GetValueDHTController) {
	api := router.Group("/api/sensor-dht11")
	{
		api.POST("/create", valueController.SaveValue)
		api.GET("/get/:user_id", getValueController.GetValue)
	}
}
