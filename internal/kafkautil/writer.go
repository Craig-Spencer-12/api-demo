package kafkautil

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func NewWriter(brokerURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(brokerURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func Produce(w *kafka.Writer, value []byte) error {
	return w.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: value,
		},
	)
}
