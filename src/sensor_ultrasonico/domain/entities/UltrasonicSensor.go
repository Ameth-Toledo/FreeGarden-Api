package entities

type UltrasonicSensor struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	Distance float64 `json:"distance_cm"`
}
