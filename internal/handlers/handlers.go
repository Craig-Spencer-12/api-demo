package handlers

import (
	"github.com/Craig-Spencer-12/api-demo/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Handler struct {
	Writer   *kafka.Writer
	Services *services.Services
}

func InitHandler(handler *gin.Engine, s services.Services, kw *kafka.Writer) {
	usersGroup := handler.Group("/users")
	NewUserRoutes(usersGroup, s.Users, kw)
}
