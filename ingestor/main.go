package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/segmentio/kafka-go"
)

type Telemetry struct {
	TruckID   string  `json:"truck_id"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
	Speed     float64 `json:"speed"`
	Timestamp float64 `json:"timestamp"`
}

func main() {

	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "kafka:9092"
	}

	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    "truck-telemetry",
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		var t Telemetry
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: Remove this
		fmt.Printf("Ingested: Truck %s at %.2f m/s\n", t.TruckID, t.Speed)

		msg, _ := json.Marshal(t)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(t.TruckID),
				Value: msg,
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
