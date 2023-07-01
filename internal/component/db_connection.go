package component

import (
	"authentication/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DatabaseConnection(cnf *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable",
		cnf.Database.Host,
		cnf.Database.Port,
		cnf.Database.User,
		cnf.Database.Password,
		cnf.Database.Name,
	)

	connection, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error when open connection %s", err.Error())
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalf("Error when ping connection %s", err.Error())
	}

	// Query to create table users
	_, err = connection.Exec(`
		DROP TABLE IF EXISTS users;

		CREATE TABLE IF NOT EXISTS users (
		id integer NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		full_name VARCHAR(255) NOT NULL,
		phone VARCHAR(255) NOT NULL,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL);

		INSERT INTO users (full_name, phone, username, password) VALUES
		('Test Full Name', '08123456789', 'test-uname', '$2a$12$bXJIkQIBLc9oIiog8llObu2wx42eB3tic3c3pZoT1SKJDbwCxhdRe')
		`) // try using method post with raw JSON = {"usename": "test-uname", "password": "12345678"}
	if err != nil {
		log.Fatalf("Error when create table users %s", err.Error())
	}

	return connection
}
