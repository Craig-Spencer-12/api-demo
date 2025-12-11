package db

import (
	"database/sql"
	"log"
)

type SQL struct {
	Pool *sql.DB
}

func New(url string) (*SQL, error) {

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("Error opening database :%v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	database := &SQL{
		Pool: db,
	}

	return database, nil
}

func (p *SQL) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
