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

	redisURL := "redis:6379" // TODO: Change to env variable
	redisDB := internal.NewRedisRepo(redisURL)
	defer redisDB.Close()

	fmt.Println("Kafka consumer started...")

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			truckIDs, err := redisDB.GetAllTruckIDs()
			if err != nil {
				fmt.Printf("failed to get truck IDs: %v\n", err)
				continue
			}

			for _, truckID := range truckIDs {
				speeds, err := redisDB.GetTruckSpeeds(truckID)
				if err != nil {
					fmt.Printf("failed to get speeds for %s: %v\n", truckID, err)
					continue
				}

				if len(speeds) == 0 {
					continue
				}

				avg := averageSpeed(speeds)
				database.AddAverageSpeed(dto.Telemetry{
					TruckID: truckID,
					Speed:   avg,
				})
				redisDB.ClearTruckSpeeds(truckID)
			}
		}
	}()

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

		err = redisDB.AddTruckSpeed(truckData)
		if err != nil {
			fmt.Printf("failed to add speed to redis: %v\n", err)
		}

		// database.AddAverageSpeed(truckData) // Replace this with posting the average every 20 seconds
		// fmt.Printf("Received event: %s\n", msg.Value)
	}
}

func averageSpeed(speeds []float64) float64 {
	if len(speeds) == 0 {
		return 0
	}
	sum := 0.0
	for _, s := range speeds {
		sum += s
	}
	return sum / float64(len(speeds))
}
