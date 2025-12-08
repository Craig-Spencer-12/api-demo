package main

import (
	"context"
	"fmt"

	"github.com/Craig-Spencer-12/api-demo/internal/kafkautil"
)

func main() {
	reader := kafkautil.NewReader(
		[]string{"localhost:9092"},
		"user-events",
		"user-events-consumer-group",
	)

	fmt.Println("Kafka consumer started...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received event: %s\n", msg.Value)
	}
}
