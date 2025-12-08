package services

import (
	"github.com/Craig-Spencer-12/api-demo/internal/services/users"
)

type Services struct {
	Users users.Service
}

// TODO: add db as a param
func NewServices() *Services {
	return &Services{
		Users: users.Service{},
	}
}
