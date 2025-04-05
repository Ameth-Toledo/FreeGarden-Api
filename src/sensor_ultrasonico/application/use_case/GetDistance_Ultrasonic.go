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

// GetLastDistanceByUserID obtiene la última distancia registrada para un usuario.
func (s *GetDistanceUltrasonic) GetLastDistanceByUserID(userID int) (float64, error) {
	// Llama al repositorio para obtener la última distancia
	return s.repo.GetLastDistanceByUserID(userID)
}

// IsContainerEmpty verifica si el recipiente está vacío según la distancia.
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
