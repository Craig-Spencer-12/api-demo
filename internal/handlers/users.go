package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Craig-Spencer-12/api-demo/internal/entity/dto"
	"github.com/Craig-Spencer-12/api-demo/internal/kafkautil"
	"github.com/Craig-Spencer-12/api-demo/internal/services/users"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type UserRoutes struct {
	u  users.Service
	kw *kafka.Writer
}

func NewUserRoutes(handler *gin.RouterGroup, u users.Service, kw *kafka.Writer) {
	r := &UserRoutes{u, kw}

	h := handler.Group("")

	h.GET("", r.GetAllUsers)
	h.GET(":id", r.GetUserByID)
	h.POST("", r.CreateUser)
}

func (ur *UserRoutes) CreateUser(c *gin.Context) {
	var userReq dto.CreateUserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	user, err := ur.u.CreateUser(userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "placeholder error"})
		return
	}

	data, _ := json.Marshal(gin.H{"message": "User Created", "user": user})
	if err := kafkautil.Produce(ur.kw, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"created user": user})
}

func (ur *UserRoutes) GetAllUsers(c *gin.Context) {
	messages, err := ur.u.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "placeholder error"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (ur *UserRoutes) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid id"})
		return
	}

	user, err := ur.u.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
