package usecases

import "example-service/internal/databases"

type Usecases struct {
	PostgresDB *databases.PostgresRepo
	RedisDB    *databases.RedisRepo
}

func NewUsecases(postgresDB *databases.PostgresRepo, redisDB *databases.RedisRepo) *Usecases {
	return &Usecases{PostgresDB: postgresDB, RedisDB: redisDB}
}
