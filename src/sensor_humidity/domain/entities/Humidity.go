package entities

import "time"

type Humidity struct {
	Id       int       `json:"id"`
	Humidity float64   `json:"humidity"`
	Date     time.Time `json:"date"`
	UserID   int       `json:"user_id"`
}
