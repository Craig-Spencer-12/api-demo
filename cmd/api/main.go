package main

import (
	"github.com/Craig-Spencer-12/api-demo/internal/handlers"
	"github.com/Craig-Spencer-12/api-demo/internal/kafkautil"
	"github.com/Craig-Spencer-12/api-demo/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	writer := kafkautil.NewWriter("localhost:9092", "user-events")

	// TODO: database creation goes here, and is input into NewServices
	service := services.NewServices()
	handlers.InitHandler(r, *service, writer)

	r.Run(":8080")
}
