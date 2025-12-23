package main

import (
	"example-service/internal/app"
	"example-service/internal/databases"
	"example-service/internal/endpoints"
	"example-service/internal/usecases"
	"fmt"
	"log"
	"os"

	"common/kafkautil"
)

func main() {
	kafkaURL := os.Getenv("KAFKA_URL")
	reader := kafkautil.NewReader(
		[]string{kafkaURL},
		"truck-telemetry",
		"truck-telemetry-group",
	)

	dbURL := os.Getenv("DB_URL")
	database, err := databases.NewPostgresRepo(dbURL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - databases.New: %w", err))
	}
	defer database.Close()

	redisURL := os.Getenv("REDIS_URL")
	redisDB := databases.NewRedisRepo(redisURL)
	defer redisDB.Close()

	usecases := usecases.NewUsecases(database, redisDB)

	endpoints := endpoints.NewEndpoints(usecases)

	fmt.Println("HTTP service started...")
	go endpoints.Foo()

	fmt.Println("Kafka consumer started...")
	app.RunConsumer(reader, usecases)
}
