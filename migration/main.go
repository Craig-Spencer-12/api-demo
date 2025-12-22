package main

import (
	"database/sql"
	"embed"
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const (
	_defaultAttempts     = 20
	_defaultTimeout      = time.Second
	_directoryPermission = 0o755
)

//go:embed all:schemas
var content embed.FS

func main() {

	url := os.Getenv("DB_URL")
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	migrationsSource, err := iofs.New(content, "schemas")
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("iofs", migrationsSource, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}

	log.Print("Success Migration")
}
