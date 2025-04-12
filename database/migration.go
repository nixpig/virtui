package database

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed sql/*_connections.*.sql
var Migrations embed.FS

type Migrator interface {
	Up() error
	Down() error
}

type Migration struct {
	migrate *migrate.Migrate
}

func (m *Migration) Up() error {
	return m.migrate.Up()
}

func (m *Migration) Down() error {
	return m.migrate.Down()
}

func NewMigration(db *sql.DB, migrations embed.FS) (*Migration, error) {
	driver, err := iofs.New(migrations, "sql")
	if err != nil {
		return nil, fmt.Errorf("create driver: %w", err)
	}

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, fmt.Errorf("sqlite3 with instance: %w", err)
	}

	m, err := migrate.NewWithInstance("file", driver, "sqlite3", instance)
	if err != nil {
		return nil, fmt.Errorf("create migration: %w", err)
	}

	return &Migration{migrate: m}, nil
}
