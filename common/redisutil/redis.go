package redisutil

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
	Ctx    context.Context
}

func New(addr, password string, db int) (*Redis, error) {
	r := &Redis{
		Ctx: context.Background(),
	}

	r.Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(r.Ctx, 5*time.Second)
	defer cancel()

	if err := r.Client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
		return nil, err
	}

	return r, nil
}

func (r *Redis) Close() {
	if r.Client != nil {
		if err := r.Client.Close(); err != nil {
			log.Printf("Error closing Redis client: %v", err)
		}
	}
}
