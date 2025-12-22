package internal

import (
	"common/dto"
	"common/postgresutil"
	"fmt"
	"strconv"
)

type Repo struct {
	*postgresutil.SQL
}

func NewPostgresRepo(url string) (*Repo, error) {
	db, err := postgresutil.New(url)
	return &Repo{SQL: db}, err
}

func (r *Repo) AddAverageSpeed(truckData dto.Telemetry) error {

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
