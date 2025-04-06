package infrastructure

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_pH/domain"
	"FreeGarden/src/sensor_pH/domain/entities"
	"fmt"
)

type MySQLpHRepository struct {
	db *core.Conn_MySQL
}

func NewMySQLpHRepository(db *core.Conn_MySQL) domain.PHRepository {
	return &MySQLpHRepository{db: db}
}

func (repo *MySQLpHRepository) Save(sensor entities.PhSensor) (entities.PhSensor, error) {
	// Definir la consulta SQL para insertar el sensor
	query := `INSERT INTO water_ph_sensors (user_id, ph_value) VALUES (?, ?)`

	// Ejecutar la consulta utilizando la conexión y los valores del sensor
	result, err := repo.db.ExecutePreparedQuery(query, sensor.UserID, sensor.PhValue)
	if err != nil {
		return sensor, fmt.Errorf("error al guardar el sensor de pH: %w", err)
	}

	// Obtener el ID del registro insertado
	id, err := result.LastInsertId()
	if err != nil {
		return sensor, fmt.Errorf("error al obtener el ID del sensor insertado: %w", err)
	}

	sensor.ID = int(id)
	return sensor, nil
}

func (repo *MySQLpHRepository) GetLastValueByUserID(userID int) (entities.PhSensor, error) {
	query := `SELECT id, user_id, ph_value FROM water_ph_sensors WHERE user_id = ? ORDER BY id DESC LIMIT 1`
	rows, err := repo.db.FetchRows(query, userID)
	if err != nil {
		return entities.PhSensor{}, fmt.Errorf("error al obtener el último valor de pH: %w", err)
	}
	defer rows.Close()

	var sensor entities.PhSensor
	if rows.Next() {
		if err := rows.Scan(&sensor.ID, &sensor.UserID, &sensor.PhValue); err != nil {
			return entities.PhSensor{}, fmt.Errorf("error al escanear los datos del sensor: %w", err)
		}
	} else {
		return entities.PhSensor{}, fmt.Errorf("no se encontró un valor de pH para el usuario %d", userID)
	}

	return sensor, nil
}
