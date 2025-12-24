package main

import (
	"common/consts"
	"example-service/internal/app"
	"example-service/internal/databases"
	"example-service/internal/endpoints"
	"example-service/internal/usecases"
	"fmt"
	"log"
	"os"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	database, err := databases.NewPostgresRepo(dbURL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - databases.New: %w", err))
	}
	defer database.Close()

	redisURL := os.Getenv("REDIS_URL")
	redisDB := databases.NewRedisRepo(redisURL) // TODO: Create and handle an error
	defer redisDB.Close()

	usecases := usecases.NewUsecases(database, redisDB)

	endpoints := endpoints.NewEndpoints(usecases)

	kafkaURL := os.Getenv("KAFKA_URL")
	for _, topic := range consts.Topics {
		fmt.Println("Kafka consumer started...")
		consumer := app.NewConsumer(kafkaURL, topic, usecases)
		go consumer.Start()
	}

	fmt.Println("HTTP service started...")
	go endpoints.Foo()

	select {} // blocks
}
