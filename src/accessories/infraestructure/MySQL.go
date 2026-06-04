package infrastructure

import (
	"fmt"
	"log"
	"proyecto/src/accessories/domain/entities"
	"proyecto/src/core"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn == nil || conn.DB == nil || conn.Err != "" { // Verificación de conexión
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Guardar accesorio
func (mysql *MySQL) Save(name, description string) error {
	query := "INSERT INTO accesories (name, description) VALUES (?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, name, description)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

// Eliminar accesorio
func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM accesories WHERE id = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún accesorio con el ID %d", id)
	}
	log.Printf("[MySQL] - Accesorio eliminado con ID: %d", id)
	return nil
}

// Ver todos los accesorios
func (mysql *MySQL) ViewAll() ([]entities.Accessory, error) {
	query := "SELECT id, name, description FROM accesories"

	rows := mysql.conn.FetchRows(query)
	if rows == nil {
		return nil, fmt.Errorf("error al recuperar los accesorios")
	}
	defer rows.Close()

	var accessories []entities.Accessory
	for rows.Next() {
		var accessory entities.Accessory
		if err := rows.Scan(&accessory.Id, &accessory.Name, &accessory.Description); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		accessories = append(accessories, accessory)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return accessories, nil
}

// Ver un accesorio por ID
func (mysql *MySQL) ViewOne(id int) (*entities.Accessory, error) {
	query := "SELECT id, name, description FROM accesories WHERE id = ?"

	rows := mysql.conn.FetchRows(query, id)
	if rows == nil {
		return nil, fmt.Errorf("error al recuperar el accesorio")
	}
	defer rows.Close()

	var accessory entities.Accessory
	if rows.Next() {
		if err := rows.Scan(&accessory.Id, &accessory.Name, &accessory.Description); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no se encontró el accesorio con ID %d", id)
	}

	return &accessory, nil
}

// Editar accesorio
func (mysql *MySQL) Edit(id int, name, description string) error {
	query := "UPDATE accesories SET name = ?, description = ? WHERE id = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, name, description, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún accesorio con el ID %d", id)
	}

	log.Printf("[MySQL] - Accesorio actualizado con ID: %d", id)
	return nil
}
