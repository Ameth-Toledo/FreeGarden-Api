package controllers

import (
	"net/http"
	"strconv"

	"FreeGarden/src/sensor_pH/application/use_case"
	"github.com/gin-gonic/gin"
)

type GetValuePhController struct {
	getValue *use_case.GetValuePH
}

func NewGetValuePhController(getValue *use_case.GetValuePH) *GetValuePhController {
	return &GetValuePhController{getValue: getValue}
}

func (controller *GetValuePhController) GetValue(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id"})
		return
	}

	value, err := strconv.Atoi(c.Param("ph_value"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ph_value"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "value": value})
}
