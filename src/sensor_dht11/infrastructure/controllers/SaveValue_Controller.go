package controllers

import (
	"FreeGarden/src/sensor_dht11/application/use_case"
	"FreeGarden/src/sensor_dht11/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SaveValueController struct {
	saveValueUseCase *use_case.CreateDHT
}

func NewSaveValueController(saveValueUseCase *use_case.CreateDHT) *SaveValueController {
	return &SaveValueController{saveValueUseCase: saveValueUseCase}
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

	c.JSON(http.StatusOK, gin.H{"message": "Value temperature saved successfully", "data": savedSensor})
}
