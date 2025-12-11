package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Craig-Spencer-12/api-demo/internal/handlers"
	"github.com/Craig-Spencer-12/api-demo/internal/kafkautil"
	"github.com/Craig-Spencer-12/api-demo/internal/services"
	"github.com/Craig-Spencer-12/api-demo/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	writer := kafkautil.NewWriter("localhost:9092", "user-events")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	database, err := db.New(dbURL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - db.New: %w", err))
	}
	defer database.Close()

	service := services.NewServices(database)
	handlers.InitHandler(r, *service, writer)

	r.Run(":8080")
}
