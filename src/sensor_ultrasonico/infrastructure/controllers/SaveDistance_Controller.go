package controller

import (
	"FreeGarden/src/sensor_ultrasonico/application/repositories"
	"FreeGarden/src/sensor_ultrasonico/application/use_case"
	"FreeGarden/src/sensor_ultrasonico/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SaveDistanceController struct {
	saveUseCase         *use_case.SaveUltrasonic
	serviceNotification *repositories.ServiceNotification
}

func NewSaveDistanceController(
	saveUseCase *use_case.SaveUltrasonic,
	serviceNotification *repositories.ServiceNotification,
) *SaveDistanceController {
	return &SaveDistanceController{
		saveUseCase:         saveUseCase,
		serviceNotification: serviceNotification,
	}
}

func (controller *SaveDistanceController) SaveDistance(c *gin.Context) {
	var sensor entities.UltrasonicSensor

	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	savedSensor, err := controller.saveUseCase.SaveDistance(sensor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving data"})
		return
	}

	// Notificar a RabbitMQ
	if err := controller.serviceNotification.NotifyAppointmentCreated(savedSensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error publishing event to RabbitMQ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Distance saved and event published successfully",
		"data":    savedSensor,
	})
}
