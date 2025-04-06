package routes

import (
	"FreeGarden/src/sensor_pH/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, valueController *controllers.SaveValueController, getValueController *controllers.GetValuePhController) {
	api := router.Group("/api/sensor-ph")
	{
		api.POST("/create", valueController.SaveValue)
		api.GET("/get:user_id", getValueController.GetValue)
	}
}
