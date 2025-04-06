package adapters_h

import (
	"database/sql"
	"fmt"

	"FreeGarden/src/sensor_humidity/domain"
	"FreeGarden/src/sensor_humidity/domain/entities"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) domain.IHumidity {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(sensor entities.Humidity) (entities.Humidity, error) {
	query := `INSERT INTO soil_moisture_sensors (humidity, user_id) VALUES (?, ?, ?)`
	_, err := m.conn.Exec(query, sensor.Humidity, sensor.Id)
	if err != nil {
		return entities.Humidity{}, fmt.Errorf("error al guardar el registro: %v", err)
	}
	return sensor, nil
}

func (r *MySQL) GetLastValueByUserID(userID int) (entities.Humidity, error) {
	var sensor entities.Humidity
	query := "SELECT id, user_id, humidity FROM soils_moisture_sensors WHERE user_id = ? ORDER BY id DESC LIMIT 1"
	err := r.conn.QueryRow(query, userID).Scan(&sensor.Id, &sensor.UserID, &sensor.Humidity)
	if err != nil {
		return entities.Humidity{}, err
	}
	return sensor, nil
}
