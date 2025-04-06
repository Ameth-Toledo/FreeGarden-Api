package adapters

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_dht11/domain/entities"
	"fmt"
)

type MySQLDHTRepository struct {
	db *core.Conn_MySQL
}

func NewMySQLDHTRepository(db *core.Conn_MySQL) *MySQLDHTRepository {
	return &MySQLDHTRepository{db: db}
}

func (r *MySQLDHTRepository) Save(DHT entities.DHT11Sensor) (entities.DHT11Sensor, error) {
	query := "INSERT INTO dht11_sensors (user_id, temperature, humidity) VALUES (?, ?, ?)"
	result, err := r.db.ExecutePreparedQuery(query, DHT.UserID, DHT.Temperature, DHT.Humidity)
	if err != nil {
		return DHT, fmt.Errorf("error al guardar el sensor DHT11: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return DHT, fmt.Errorf("error al obtener el ID del sensor insertado: %w", err)
	}
	DHT.ID = int(id)

	return DHT, nil
}

func (r *MySQLDHTRepository) GetLastValueByUserID(userID int) (entities.DHT11Sensor, error) {
	var sensor entities.DHT11Sensor

	query := "SELECT id, user_id, temperature, humidity FROM dht11_sensors WHERE user_id = ? ORDER BY id DESC LIMIT 1"
	rows, err := r.db.FetchRows(query, userID)
	if err != nil {
		return sensor, fmt.Errorf("error al obtener la última lectura DHT11: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&sensor.ID, &sensor.UserID, &sensor.Temperature, &sensor.Humidity)
		if err != nil {
			return sensor, fmt.Errorf("error al escanear la fila del sensor DHT11: %w", err)
		}
	} else {
		return sensor, fmt.Errorf("no se encontraron datos del sensor DHT11 para el usuario %d", userID)
	}

	return sensor, nil
}
