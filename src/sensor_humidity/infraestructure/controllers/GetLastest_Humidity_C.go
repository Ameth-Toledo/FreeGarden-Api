package controllers

import (
	"FreeGarden/src/sensor_humidity/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetLatestHumidityController struct {
	useCase *use_case.Get_Latest
}

func NewGetLatestHumidityController(useCase *use_case.Get_Latest) *GetLatestHumidityController {
	return &GetLatestHumidityController{useCase: useCase}
}

func (gl *GetLatestHumidityController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	data, err := gl.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la última medición", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
