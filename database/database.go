package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewConnection(filename string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("file:%s", filename)

	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, fmt.Errorf(
			"open database connection (%s): %w",
			connectionString, err,
		)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}
