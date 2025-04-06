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

func (repo *MySQLpHRepository) Save(ph entities.PhSensor) (entities.PhSensor, error) {

	query := `INSERT INTO ph_sensors (user_id, ph_value) VALUES (?, ?)`

	result, err := repo.db.ExecutePreparedQuery(query, ph.UserID, ph.PhValue)
	if err != nil {
		return ph, fmt.Errorf("error al guardar el ph sensor: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return ph, fmt.Errorf("error al guardar el ph sensor: %w", err)
	}
	ph.ID = int(id)
	return ph, nil
}

func (repo *MySQLpHRepository) GetLastValue(userID int) (float64, error) {
	query := `SELECT ph_value FROM ph_sensors WHERE user_id = ?`

	rows, err := repo.db.FetchRows(query, userID)
	if err != nil {
		return 0, fmt.Errorf("error al obtene la ultima lectura del ph sensor: %w", err)
	}
	defer rows.Close()

	var ph_value float64
	if rows.Next() {
		if err := rows.Scan(&ph_value); err != nil {
			return 0, fmt.Errorf("error al obtener el valor del ph sensor: %w", err)
		}
	}

	if ph_value == 0 {
		return 0, fmt.Errorf("no se encontró un sensor de ph para el usuario %d", userID)
	}

	return ph_value, nil
}
