package db

import (
	"database/sql"
	"log"
)

type SQL struct {
	Pool *sql.DB
}

func New(url string) (*SQL, error) {

	db := &SQL{}
	var err error

	db.Pool, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("Error opening database :%v", err)
	}

	if err := db.Pool.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return db, nil
}

func (p *SQL) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
