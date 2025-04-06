package controllers

import (
	"FreeGarden/src/sensor_humidity/application/use_case"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteHumidityController struct {
	useCase *use_case.Delete_Humidty
}

func NewDeleteHumiditypController(useCase *use_case.Delete_Humidty) *DeleteHumidityController {
	return &DeleteHumidityController{useCase: useCase}
}

func (ct *DeleteHumidityController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userIDStr := ctx.Param("user_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID inválido"})
		return
	}

	if err := ct.useCase.Execute(id, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar los datos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Datos eliminados correctamente"})
}
