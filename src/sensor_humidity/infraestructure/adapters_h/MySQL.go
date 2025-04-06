package adapters_h

import (
	"database/sql"
	"errors"
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
	query := `INSERT INTO humidity (humidity, date, user_id) VALUES (?, ?, ?)`
	_, err := m.conn.Exec(query, sensor.Humidity, sensor.Date, sensor.Id)
	if err != nil {
		return entities.Humidity{}, fmt.Errorf("error al guardar el registro: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetMeasurementByID(id int, userID int) (entities.Humidity, error) {
	var sensor entities.Humidity
	query := `SELECT id, humidity, date FROM humidity WHERE id = ? AND user_id = ?`
	err := m.conn.QueryRow(query, id, userID).Scan(&sensor.Id, &sensor.Humidity, &sensor.Date)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sensor, errors.New("registro no encontrado")
		}
		return sensor, fmt.Errorf("error al obtener la medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetLatestMeasurement(userID int) (entities.Humidity, error) {
	var sensor entities.Humidity
	query := `SELECT id, humidity, date FROM humidity WHERE user_id = ? ORDER BY id DESC LIMIT 1`
	err := m.conn.QueryRow(query, userID).Scan(&sensor.Id, &sensor.Humidity, &sensor.Date)
	if err != nil {
		return sensor, fmt.Errorf("error al obtener la última medición: %v", err)
	}
	return sensor, nil
}

func (m *MySQL) GetAllMeasurements(userID int) ([]entities.Humidity, error) {
	query := "SELECT id, humidity, date FROM humidity WHERE user_id = ?"
	rows, err := m.conn.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los registros: %v", err)
	}
	defer rows.Close()

	var records []entities.Humidity
	for rows.Next() {
		var sensor entities.Humidity
		if err := rows.Scan(&sensor.Id, &sensor.Humidity, &sensor.Date); err != nil {
			return nil, fmt.Errorf("error al escanear los registros: %v", err)
		}
		records = append(records, sensor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al recorrer los registros: %v", err)
	}

	return records, nil
}

func (m *MySQL) DeleteMeasurement(id int, userID int) error {
	query := "DELETE FROM humidity WHERE id = ? AND user_id = ?"
	_, err := m.conn.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error al eliminar el registro: %v", err)
	}
	return nil
}

func (m *MySQL) GetAverageHumidity(userID int) (float64, error) {
	var avgHumidity float64
	query := "SELECT AVG(humidity) FROM humidity WHERE user_id = ?"
	err := m.conn.QueryRow(query, userID).Scan(&avgHumidity)
	if err != nil {
		return 0, fmt.Errorf("error al calcular el promedio de humedad: %v", err)
	}
	return avgHumidity, nil
}
