package services

import (
	"github.com/Craig-Spencer-12/api-demo/internal/services/users"
	"github.com/Craig-Spencer-12/api-demo/pkg/db"
)

type Services struct {
	Users users.Service
}

func NewServices(db *db.SQL) *Services {
	return &Services{
		Users: users.NewService(db),
	}
}
