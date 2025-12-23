package app

import (
	"common/dto"
	"context"
	"encoding/json"
	"example-service/internal/usecases"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func RunConsumer(reader *kafka.Reader, usecases *usecases.Usecases) {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			usecases.CalculateAverageTruckSpeeds()
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
			continue
		}

		err = usecases.RedisDB.AddTruckSpeed(truckData)
		if err != nil {
			fmt.Printf("failed to add speed to redis: %v\n", err)
		}
	}
}
