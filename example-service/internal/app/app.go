package app

import (
	"common/consts"
	"common/dto"
	"context"
	"encoding/json"
	"example-service/internal/usecases"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	r  *kafka.Reader
	uc *usecases.Usecases
}

func NewConsumer(kafkaURL string, topic consts.TopicConfig, uc *usecases.Usecases) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},
		Topic:   topic.Name,
		GroupID: topic.ConsumerGroup,
	})

	return &Consumer{
		r:  reader,
		uc: uc,
	}
}

func (c *Consumer) Start() {
	switch c.r.Config().Topic {
	case consts.TruckTelemetry:
		c.runTruckTelemetryConsumer()
	case consts.SpeedAlert:
		c.runSpeedAlertConsumer()
	case consts.AccelAlert:
		c.runAccelAlertConsumer()
	case consts.EndTrip:
		c.runEndTripConsumer()
	default:
		fmt.Println("invalid topic")
	}
}

func (c *Consumer) runTruckTelemetryConsumer() {
	for {
		msg, err := c.r.ReadMessage(context.Background())
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

		err = c.uc.RedisDB.AddTruckSpeed(truckData) // TODO: consumer should not talk directly to redis
		if err != nil {
			fmt.Printf("failed to add speed to redis: %v\n", err)
		}
	}
}

func (c *Consumer) runSpeedAlertConsumer() {

}

func (c *Consumer) runAccelAlertConsumer() {

}

func (c *Consumer) runEndTripConsumer() {

}

// Example ticker logic
// ticker := time.NewTicker(20 * time.Second)
// 	defer ticker.Stop()

// 	go func() {
// 		for range ticker.C {
// 			usecases.CalculateAverageTruckSpeeds()
// 		}
// 	}()
