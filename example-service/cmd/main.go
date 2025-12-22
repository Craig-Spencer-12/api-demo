package main

import (
	"context"
	"encoding/json"
	"example-service/internal"
	"fmt"
	"log"
	"os"
	"time"

	"common/dto"
	"common/kafkautil"
)

func main() {
	reader := kafkautil.NewReader(
		[]string{"kafka:9092"},
		"truck-telemetry",
		"truck-telemetry-group",
	)

	dbURL := os.Getenv("DB_URL")
	database, err := internal.NewPostgresRepo(dbURL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - db.New: %w", err))
	}
	defer database.Close()

	fmt.Println("Kafka consumer started...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("fetching message: %v, retrying in 1s...\n", err)
			time.Sleep(time.Second)
			continue
		}

		var truckData dto.Telemetry

		err = json.Unmarshal(msg.Value, &truckData)
		if err != nil {
			fmt.Printf("invalid telemetry message: %v\n", err)
			continue // skip bad messages
		}

		database.Create(truckData)
		fmt.Printf("Received event: %s\n", msg.Value)
	}
}
