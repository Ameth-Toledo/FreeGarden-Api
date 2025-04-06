package routes_h

import (
	"FreeGarden/src/core/security"
	"FreeGarden/src/sensor_humidity/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterHumidityRoutes(
	router *gin.Engine,
	createController *controllers.Create_Humidity_C,
	getHumidityByIDController *controllers.GetHumidityByIDController,
	getAllController *controllers.GetAllHumidityController,
	deleteController *controllers.DeleteHumidityController,
	getAverageHumidityController *controllers.GetAverageHumidityController,
	getLatestMeasurementController *controllers.GetLatestHumidityController,
) {
	api := router.Group("/api/sensor-humidity")
	{
		api.POST("/create", createController.Execute)
		api.GET("/humidity/:id/:user_id", security.JWTMiddleware(), getHumidityByIDController.Execute)
		api.GET("/all/:user_id", security.JWTMiddleware(), getAllController.Execute)
		api.DELETE("/delete/:id/:user_id", security.JWTMiddleware(), deleteController.Execute)
		api.GET("/humidity/average/:user_id", security.JWTMiddleware(), getAverageHumidityController.Execute)
		api.GET("/humidity/latest/:user_id", security.JWTMiddleware(), getLatestMeasurementController.Execute)
	}
}
