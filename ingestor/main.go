package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"common/consts"
	"common/dto"

	"github.com/segmentio/kafka-go"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")

	writers := make(map[string]*kafka.Writer)
	for _, topic := range consts.Topics {
		writers[topic.Name] = &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    topic.Name,
			Balancer: &kafka.LeastBytes{},
		}
		defer writers[topic.Name].Close()
	}

	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		topic := r.Header.Get("X-Topic")
		if topic == "" {
			http.Error(w, "Missing X-Topic header", http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var meta dto.Meta
		if err := json.Unmarshal(body, &meta); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = writers[topic].WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(meta.TruckID),
				Value: body,
			},
		)
		if err != nil {
			log.Println("Kafka Error:", err)
		}

		w.WriteHeader(http.StatusAccepted)
	})

	fmt.Println("Ingestor listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
