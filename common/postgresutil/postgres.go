package postgresutil

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type SQL struct {
	Pool *sql.DB
}

func New(url string) (*SQL, error) {

	db := &SQL{}
	var err error

	db.Pool, err = sql.Open("pgx", url)
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
