package controller

import (
	"FreeGarden/src/sensor_ultrasonico/application/use_case"
	"FreeGarden/src/sensor_ultrasonico/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SaveDistanceController maneja las solicitudes de guardar las distancias de los sensores.
type SaveDistanceController struct {
	saveUseCase *use_case.SaveUltrasonic
}

// NewSaveDistanceController crea una nueva instancia de SaveDistanceController.
func NewSaveDistanceController(saveUseCase *use_case.SaveUltrasonic) *SaveDistanceController {
	return &SaveDistanceController{saveUseCase: saveUseCase}
}

// SaveDistance maneja la solicitud HTTP para guardar una nueva distancia.
func (controller *SaveDistanceController) SaveDistance(c *gin.Context) {
	var sensor entities.UltrasonicSensor

	// Bindear el cuerpo JSON de la solicitud a la entidad UltrasonicSensor
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Llamar al caso de uso para guardar la distancia
	savedSensor, err := controller.saveUseCase.SaveDistance(sensor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Distance saved successfully", "data": savedSensor})
}
