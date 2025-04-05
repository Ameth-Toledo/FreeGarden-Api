package use_case

import (
	"FreeGarden/src/sensor_ultrasonico/domain"
)

type GetDistanceUltrasonic struct {
	repo domain.UltrasonicRepository
}

func NewUltrasonicGetDistance(repo domain.UltrasonicRepository) *GetDistanceUltrasonic {
	return &GetDistanceUltrasonic{repo: repo}
}

func (s *GetDistanceUltrasonic) IsContainerEmpty(userID int, emptyThreshold float64) (bool, error) {
	lastDistance, err := s.repo.GetLastDistanceByUserID(userID)
	if err != nil {
		return false, err
	}
	if lastDistance >= emptyThreshold {
		return true, nil
	}
	return false, nil
}
