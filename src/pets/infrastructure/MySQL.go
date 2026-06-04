package infrastructure

import (
	"fmt"
	"log"
	"proyecto/src/core"
	"proyecto/src/pets/domain/entities"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(name, raza string) error {
	query := "INSERT INTO pets (name, raza) VALUES (?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, name, raza)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM pets WHERE id = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ninguna mascota con el ID %d", id)
	}
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) ViewAll() ([]entities.Pet, error) {
	query := "SELECT * FROM pets"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var pets []entities.Pet
	for rows.Next() {
		var pet entities.Pet
		if err := rows.Scan(&pet.Id, &pet.Name, &pet.Raza); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		pets = append(pets, pet)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return pets, nil
}

func (mysql *MySQL) ViewOne(id int) (*entities.Pet, error) {
	query := "SELECT * FROM pets WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var pet entities.Pet
	if rows.Next() {
		if err := rows.Scan(&pet.Id, &pet.Name, &pet.Raza); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no se encontró ninguna mascota con el ID %d", id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}

	return &pet, nil
}
func (mysql *MySQL) Edit(id int, name, raza string) error {
	query := "UPDATE pets SET name = ?, raza = ? WHERE id = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, name, raza, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	new_id := id
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ninguna mascota con el ID %d", new_id)
	}

	return nil
}
