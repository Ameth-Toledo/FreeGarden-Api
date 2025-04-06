package routes

import (
	"FreeGarden/src/sensor_pH/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, valueController *controllers.SaveValueController, getValueController *controllers.GetValuePhController) {
	api := router.Group("/sensor")
	{
		api.POST("/ph", valueController.SaveValue)
		api.GET("/ph", getValueController.GetValue)
	}
}
