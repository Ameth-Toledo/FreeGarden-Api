package controllers

import (
	"net/http"
	"strconv"

	"FreeGarden/src/sensor_humidity/application/use_case"
	"github.com/gin-gonic/gin"
)

type GetValueHumidityController struct {
	getHumidity *use_case.GetValueHumidity
}

func NewGetValueHUmidityController(getHumidity *use_case.GetValueHumidity) *GetValueHumidityController {
	return &GetValueHumidityController{getHumidity: getHumidity}
}

func (controller *GetValueHumidityController) GetValue(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id"})
		return
	}

	value, err := strconv.Atoi(c.Param("humidity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid humidity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "value": value})
}
