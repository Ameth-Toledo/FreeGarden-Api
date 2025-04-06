package controllers

import (
	"FreeGarden/src/sensor_humidity/application/repositories"
	"FreeGarden/src/sensor_humidity/application/use_case"
	"FreeGarden/src/sensor_humidity/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SaveHumidityController struct {
	saveHumidityUseCase *use_case.SaveHumidity
	serviceNotification *repositories.ServiceNotification
}

func NewSaveHumidityController(saveHumidityUseCase *use_case.SaveHumidity, serviceNotification *repositories.ServiceNotification) *SaveHumidityController {
	return &SaveHumidityController{
		saveHumidityUseCase: saveHumidityUseCase,
		serviceNotification: serviceNotification,
	}
}

func (controller *SaveHumidityController) SaveHumidity(c *gin.Context) {
	var sensor entities.Humidity

	// Recepción de datos del ESP32 en formato JSON
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Guardar el valor de humedad en la base de datos
	savedSensor, err := controller.saveHumidityUseCase.SaveValue(sensor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving data"})
		return
	}

	// Notificar a RabbitMQ después de guardar los datos
	if err := controller.serviceNotification.NotifyAppointmentCreated(savedSensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error publishing event to RabbitMQ"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Humidity value saved and event published successfully",
		"data":    savedSensor,
	})
}
