package entities

type PhSensor struct {
	ID      int     `json:"id"`
	UserID  int     `json:"user_id"`
	PhValue float64 `json:"ph_value"`
}
