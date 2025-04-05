package infrastructure

import (
	"FreeGarden/src/core"
	"FreeGarden/src/sensor_ultrasonico/domain"
	"FreeGarden/src/sensor_ultrasonico/domain/entities"
	"fmt"
)

type MySQLUltrasonicRepository struct {
	db *core.Conn_MySQL
}

// NewMySQLUltrasonicRepository crea un nuevo repositorio para sensores ultrasónicos.
func NewMySQLUltrasonicRepository(db *core.Conn_MySQL) domain.UltrasonicRepository {
	return &MySQLUltrasonicRepository{db: db}
}

// Save guarda un sensor ultrasónico en la base de datos.
func (repo *MySQLUltrasonicRepository) Save(ultrasonic entities.UltrasonicSensor) (entities.UltrasonicSensor, error) {
	// Definir la consulta SQL para insertar el sensor
	query := `INSERT INTO ultrasonic_sensors (user_id, distance) VALUES (?, ?)`

	// Ejecutar la consulta utilizando la conexión y los valores del sensor
	result, err := repo.db.ExecutePreparedQuery(query, ultrasonic.UserID, ultrasonic.Distance)
	if err != nil {
		return ultrasonic, fmt.Errorf("error al guardar el sensor ultrasónico: %w", err)
	}

	// Obtener el ID del registro insertado
	id, err := result.LastInsertId()
	if err != nil {
		return ultrasonic, fmt.Errorf("error al obtener el ID del sensor insertado: %w", err)
	}

	// Asignar el ID insertado al objeto UltrasonicSensor y devolverlo
	ultrasonic.ID = int(id)
	return ultrasonic, nil
}

// GetLastDistanceByUserID obtiene la última distancia registrada para un usuario.
func (repo *MySQLUltrasonicRepository) GetLastDistanceByUserID(userID int) (float64, error) {
	query := `SELECT distance FROM ultrasonic_sensors WHERE user_id = ? ORDER BY id DESC LIMIT 1`
	rows, err := repo.db.FetchRows(query, userID)
	if err != nil {
		return 0, fmt.Errorf("error al obtener la última distancia: %w", err)
	}
	defer rows.Close()

	var distance float64
	if rows.Next() {
		if err := rows.Scan(&distance); err != nil {
			return 0, fmt.Errorf("error al leer la distancia del sensor: %w", err)
		}
	}

	// Si no se encuentra una distancia, devolver 0
	if distance == 0 {
		return 0, fmt.Errorf("no se encontró una distancia para el usuario %d", userID)
	}

	return distance, nil
}
