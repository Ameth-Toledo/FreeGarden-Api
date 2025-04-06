package dependencies_h

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_humidity/application/repositories"
	"FreeGarden/src/sensor_humidity/application/use_case"
	"FreeGarden/src/sensor_humidity/infraestructure/adapters_h"
	"FreeGarden/src/sensor_humidity/infraestructure/controllers"
	"log"
)

func Init(pool *core.Conn_MySQL) (
	*controllers.Create_Humidity_C,
	*controllers.GetHumidityByIDController,
	*controllers.GetAllHumidityController,
	*controllers.DeleteHumidityController,
	*controllers.GetAverageHumidityController,
	*controllers.GetLatestHumidityController,
	*repositories.ServiceNotification,

	error,
) {

	repository := adapters_h.NewMySQL(pool.DB)

	rabbitMQAdapter, err := adapters_h.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)

	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	createUseCase := use_case.NewCreate_Humidity(repository, serviceNotification)
	getHeartRateByIDUseCase := use_case.NewGetMeasurementByID(repository)
	getAllUseCase := use_case.NewGet_All(repository)
	deleteUseCase := use_case.NewDelete_Humidty(repository)
	getAverageHeartRateUseCase := use_case.NewGetAverageHumidity(repository)
	getLatestMeasurementUseCase := use_case.NewGet_Latest(repository)

	createController := controllers.NewCreate_Humidity_C(createUseCase)
	getHeartRateByIDController := controllers.NewGetHumidityIDController(getHeartRateByIDUseCase)
	getAllController := controllers.NewGetAllHumidityController(getAllUseCase)
	deleteController := controllers.NewDeleteHumiditypController(deleteUseCase)
	getAverageHeartRateController := controllers.NewGetAverageHumidityController(getAverageHeartRateUseCase)
	getLatestMeasurementController := controllers.NewGetLatestHumidityController(getLatestMeasurementUseCase)

	return createController,
		getHeartRateByIDController,
		getAllController,
		deleteController,
		getAverageHeartRateController,
		getLatestMeasurementController,
		serviceNotification,
		nil
}
