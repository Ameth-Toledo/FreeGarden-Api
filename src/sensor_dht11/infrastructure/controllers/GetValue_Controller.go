package controllers

import (
	"net/http"
	"strconv"

	"FreeGarden/src/sensor_dht11/application/use_case"
	"github.com/gin-gonic/gin"
)

type GetValueDHTController struct {
	getValue *use_case.GetValueDHT
}

func NewGetValueDHTController(getValue *use_case.GetValueDHT) *GetValueDHTController {
	return &GetValueDHTController{getValue: getValue}
}

func (controller *GetValueDHTController) GetValue(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id"})
		return
	}

	value, err := strconv.Atoi(c.Param("temperature"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid temperature"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "value": value})
}
