package controllers

import (
	"FreeGarden/src/sensor_humidity/application/use_case"
	"FreeGarden/src/sensor_humidity/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetHumidityByIDController struct {
	useCase *use_case.GetMeasurementByID
}

func NewGetHumidityIDController(useCase *use_case.GetMeasurementByID) *GetHumidityByIDController {
	return &GetHumidityByIDController{useCase: useCase}
}

func (gh *GetHumidityByIDController) Execute(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userIDStr := ctx.Param("user_id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	measurement, err := gh.useCase.Execute(id, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la medición", "details": err.Error()})
		return
	}
	if (measurement == entities.Humidity{}) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontró el registro con el ID proporcionado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": measurement})
}
