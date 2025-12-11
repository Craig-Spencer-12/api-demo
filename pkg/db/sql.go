package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SQL struct {
	Pool *pgxpool.Pool
}

func New(url string) (*SQL, error) {

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	runMigrations()

	db := &SQL{
		Pool: pool,
	}

	return db, nil
}

func (p *SQL) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func runMigrations(pool *pgxpool.Pool) {
	sqlBytes, err := os.ReadFile("internal/schema.sql")
	if err != nil {
		log.Fatalf("Error reading schema.sql: %v", err)
	}

	_, err = pool.Exec(context.Background(), string(sqlBytes))
	if err != nil {
		log.Fatalf("Error running schema.sql: %v", err)
	}

	log.Println("Database tables ready ✔")
}
