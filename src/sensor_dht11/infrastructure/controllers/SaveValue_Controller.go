package controllers

import (
	"FreeGarden/src/sensor_dht11/application/repositories"
	"FreeGarden/src/sensor_dht11/application/use_case"
	"FreeGarden/src/sensor_dht11/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SaveValueController struct {
	saveValueUseCase    *use_case.CreateDHT
	serviceNotification *repositories.ServiceNotification
}

func NewSaveValueController(saveValueUseCase *use_case.CreateDHT, serviceNotification *repositories.ServiceNotification) *SaveValueController {
	return &SaveValueController{
		saveValueUseCase:    saveValueUseCase,
		serviceNotification: serviceNotification,
	}
}

func (controller *SaveValueController) SaveValue(c *gin.Context) {
	var sensor entities.DHT11Sensor

	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	savedSensor, err := controller.saveValueUseCase.SaveValue(sensor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Notificar evento a RabbitMQ
	if err := controller.serviceNotification.NotifyAppoinmentCreated(savedSensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al publicar el evento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Value saved and event published successfully",
		"data":    savedSensor,
	})
}
