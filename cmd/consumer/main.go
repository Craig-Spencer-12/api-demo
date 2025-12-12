package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Craig-Spencer-12/api-demo/internal/kafkautil"
)

func main() {
	reader := kafkautil.NewReader(
		[]string{"kafka:9092"},
		"user-events",
		"user-events-consumer-group",
	)

	fmt.Println("Kafka consumer started...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("fetching message: %v, retrying in 1s...\n", err)
			time.Sleep(time.Second)
			continue
		}
		fmt.Printf("Received event: %s\n", msg.Value)
	}
}
