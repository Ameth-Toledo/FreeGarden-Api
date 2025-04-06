package dependencies

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_pH/application/use_case"
	"FreeGarden/src/sensor_pH/infrastructure"
	"FreeGarden/src/sensor_pH/infrastructure/controllers"
	"FreeGarden/src/sensor_pH/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func InitializeSensorPhDependencies() (*gin.Engine, *use_case.SavePH, *use_case.GetValuePH) {
	dbConn := core.GetDBPool()

	phRepo := infrastructure.NewMySQLpHRepository(dbConn)

	savePhUseCase := use_case.NewSavePH(phRepo)
	getValueUseCase := use_case.NewGetValuePH(phRepo)

	savePhController := controllers.NewSaveValueController(savePhUseCase)
	getValueController := controllers.NewGetValuePhController(getValueUseCase)

	router := gin.Default()

	routes.SetupRoutes(router, savePhController, getValueController)

	return router, savePhUseCase, getValueUseCase
}
