package app

import (
	"database/sql"
	"embed"
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const (
	_defaultAttempts     = 20
	_defaultTimeout      = time.Second
	_directoryPermission = 0o755
)

//go:embed all:migrations
var content embed.FS

func InitDB(url string) error {

	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	migrationsSource, err := iofs.New(content, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", migrationsSource, "postgres", driver)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}

	return nil
}
