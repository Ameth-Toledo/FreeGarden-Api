package entities

type DTH11Sensor struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
