package domain

import (
	"FreeGarden/src/sensor_ultrasonico/domain/entities"
)

type UltrasonicRepository interface {
	Save(ultrasonic entities.UltrasonicSensor) (entities.UltrasonicSensor, error)
	GetLastDistanceByUserID(userID int) (float64, error)
}
