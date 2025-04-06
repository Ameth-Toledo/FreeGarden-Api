package entities

type Humidity struct {
	Id       int     `json:"id"`
	Humidity float64 `json:"humidity"`
	UserID   int     `json:"user_id"`
}
