package controllers

import (
	"FreeGarden/src/sensor_humidity/application/use_case"
	"FreeGarden/src/sensor_humidity/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Create_Humidity_C struct {
	UseCase *use_case.CreateHumidity
}

func NewCreate_Humidity_C(useCase *use_case.CreateHumidity) *Create_Humidity_C {
	return &Create_Humidity_C{UseCase: useCase}
}

func (c *Create_Humidity_C) Execute(ctx *gin.Context) {
	var sensor entities.Humidity

	if err := ctx.ShouldBindJSON(&sensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	createdSensor, err := c.UseCase.Execute(sensor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente", "data": createdSensor})
}
