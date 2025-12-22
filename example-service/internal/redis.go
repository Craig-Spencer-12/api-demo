package internal

import (
	"common/dto"
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	*redis.Client
}

func NewRedisRepo(url string) *RedisRepo {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})

	return &RedisRepo{Client: rdb}
}

func (r *RedisRepo) AddTruckSpeed(truckData dto.Telemetry) error {
	ctx := context.Background()
	key := fmt.Sprintf("truck:%s:speeds", truckData.TruckID)
	return r.RPush(ctx, key, strconv.FormatFloat(truckData.Speed, 'f', 2, 64)).Err()
}

func (r *RedisRepo) GetAllTruckIDs() ([]string, error) {
	ctx := context.Background()
	keys, err := r.Keys(ctx, "truck:*:speeds").Result()
	if err != nil {
		return nil, err
	}

	// truckIDs := make([]string, 0, len(keys))
	// for _, key := range keys { //TODO: fix, Needlessly complicated because it could just use the key
	// 	// key format: "truck:<truckID>:speeds"
	// 	var truckID string
	// 	fmt.Sscanf(key, "truck:%s:speeds", &truckID)
	// 	truckIDs = append(truckIDs, truckID)
	// }

	return keys, nil
}

func (r *RedisRepo) GetTruckSpeeds(key string) ([]float64, error) {
	ctx := context.Background()
	vals, err := r.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	speeds := make([]float64, 0, len(vals))
	for _, v := range vals {
		speed, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue // skip invalid values
		}
		speeds = append(speeds, speed)
	}
	return speeds, nil
}

func (r *RedisRepo) ClearTruckSpeeds(truckID string) error {
	ctx := context.Background()
	key := fmt.Sprintf("truck:%s:speeds", truckID)
	return r.Del(ctx, key).Err()
}
