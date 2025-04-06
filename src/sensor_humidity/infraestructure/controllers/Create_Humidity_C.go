package controllers

import (
	"FreeGarden/src/sensor_humidity/application/use_case"
	"FreeGarden/src/sensor_humidity/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SaveHumidityController struct {
	saveHumidityUseCase *use_case.SaveHumidity
}

func NewSaveHumidityController(saveHumidityUseCase *use_case.SaveHumidity) *SaveHumidityController {
	return &SaveHumidityController{saveHumidityUseCase: saveHumidityUseCase}
}

func (controller *SaveHumidityController) SaveHumidity(c *gin.Context) {
	var sensor entities.Humidity

	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	savedSensor, err := controller.saveHumidityUseCase.SaveValue(sensor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Value Humidity saved successfully", "data": savedSensor})
}
