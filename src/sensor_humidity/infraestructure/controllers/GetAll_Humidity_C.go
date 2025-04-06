package controllers

import (
	"FreeGarden/src/sensor_humidity/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAllHumidityController struct {
	useCase *use_case.Get_All_Humidity
}

func NewGetAllHumidityController(useCase *use_case.Get_All_Humidity) *GetAllHumidityController {
	return &GetAllHumidityController{useCase: useCase}
}

func (gt *GetAllHumidityController) Execute(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	data, err := gt.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los registros", "details": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
