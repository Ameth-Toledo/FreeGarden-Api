package domain

import "FreeGarden/src/sensor_pH/domain/entities"

type PHRepository interface {
	Save(pH entities.PhSensor) (entities.PhSensor, error)
	GetLastValueByUserID(userID int) (entities.PhSensor, error)
}
