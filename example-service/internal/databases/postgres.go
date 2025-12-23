package databases

import (
	"common/dto"
	"common/postgresutil"
	"fmt"
	"strconv"
)

type PostgresRepo struct {
	*postgresutil.SQL
}

func NewPostgresRepo(url string) (*PostgresRepo, error) {
	db, err := postgresutil.New(url)
	return &PostgresRepo{SQL: db}, err
}

func (r *PostgresRepo) AddAverageSpeed(truckData dto.Telemetry) error {

	fmt.Println("Trying to enter truck data", truckData.TruckID, truckData.Speed)
	query := `
        INSERT INTO speeds (truck_id, speed)
        VALUES ($1, $2)
        RETURNING id, truck_id, speed, created_at;
    `
	_, err := r.Pool.Exec(
		query,
		truckData.TruckID,
		strconv.FormatFloat(truckData.Speed, 'f', 2, 64),
	)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
