package controller

import (
	"net/http"
	"strconv"

	"FreeGarden/src/sensor_ultrasonico/application/use_case"
	"github.com/gin-gonic/gin"
)

// GetDistanceController maneja las solicitudes de obtener la distancia del sensor.
type GetDistanceController struct {
	getDistanceUseCase *use_case.GetDistanceUltrasonic
}

// NewGetDistanceController crea una nueva instancia de GetDistanceController.
func NewGetDistanceController(getDistanceUseCase *use_case.GetDistanceUltrasonic) *GetDistanceController {
	return &GetDistanceController{getDistanceUseCase: getDistanceUseCase}
}

// GetDistance maneja la solicitud HTTP para obtener la distancia del sensor.
func (controller *GetDistanceController) GetDistance(c *gin.Context) {
	// Obtener el userID de los parámetros de la URL
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Obtener la distancia desde el caso de uso
	distance, err := controller.getDistanceUseCase.GetLastDistanceByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "last_distance": distance})
}
