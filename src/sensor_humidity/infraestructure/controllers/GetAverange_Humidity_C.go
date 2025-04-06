package controllers

import (
	"FreeGarden/src/sensor_humidity/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAverageHumidityController struct {
	useCase *use_case.GetAverageHumidity
}

func NewGetAverageHumidityController(useCase *use_case.GetAverageHumidity) *GetAverageHumidityController {
	return &GetAverageHumidityController{useCase: useCase}
}

func (gt *GetAverageHumidityController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	averageHeartRate, err := gt.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al calcular el promedio de ritmo cardíaco",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"average_heart_rate": averageHeartRate})
}
