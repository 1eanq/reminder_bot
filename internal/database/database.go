package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	id           int64
	first_name   string
	last_name    string
	username     string
	phone_number string
	country      string
}

func Connect(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	Init(db)

	return db
}

func Init(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		username VARCHAR(100),
		phone_number VARCHAR(20),
		country VARCHAR(100)
	);`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
