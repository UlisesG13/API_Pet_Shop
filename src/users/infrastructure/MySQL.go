package infrastructure

import (
	"fmt"
	"log"
	"proyecto/src/core"
	"proyecto/src/users/domain/entities"
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

func (mysql *MySQL) Register(email, password, name string) error {
	query := "INSERT INTO users (email, password, name) VALUES (?, ?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, email, password, name)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) Login(email, password string) (*entities.User, error) {
	query := "SELECT id, email, password, name FROM users WHERE email = ?"

	rows := mysql.conn.FetchRows(query, email)
	defer rows.Close()

	if rows.Next() {
		var user entities.User
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name)
		if err != nil {
			return nil, fmt.Errorf("error al escanear el usuario: %w", err)
		}
		return &user, nil
	}

	return nil, fmt.Errorf("usuario no encontrado")
}

func (mysql *MySQL) GetByEmail(email string) (*entities.User, error) {
	query := "SELECT id, email, password, name FROM users WHERE email = ?"

	rows := mysql.conn.FetchRows(query, email)
	defer rows.Close()

	if rows.Next() {
		var user entities.User
		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name)
		if err != nil {
			return nil, fmt.Errorf("error al escanear el usuario: %w", err)
		}
		return &user, nil
	}

	return nil, fmt.Errorf("usuario no encontrado")
}
