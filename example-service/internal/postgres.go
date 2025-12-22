package internal

import (
	"common/dto"
	"common/postgresutil"
)

type Repo struct {
	*postgresutil.SQL
}

func NewPostgresRepo(url string) (*Repo, error) {
	db, err := postgresutil.New(url)
	return &Repo{SQL: db}, err
}

func (r *Repo) Create(truckData dto.Telemetry) error {

	query := `
        INSERT INTO users (username, email)
        VALUES ($1, $2)
        RETURNING id, username, email, created_at;
    `
	_, err := r.Pool.Exec(
		query,
		truckData.TruckID,
		truckData.TruckID,
	)

	if err != nil {
		return err
	}

	return nil
}

// func (r *UsersRepo) GetAll() ([]dto.User, error) {
// 	query := `
//         SELECT id, username, email, created_at
//         FROM users
//         ORDER BY id ASC;
//     `
// 	rows, err := r.Pool.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	users := []dto.User{}

// 	for rows.Next() {
// 		var u dto.User
// 		err := rows.Scan(
// 			&u.ID,
// 			&u.Username,
// 			&u.Email,
// 			&u.CreatedAt,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		users = append(users, u)
// 	}

// 	return users, nil
// }

// func (r *UsersRepo) GetByID(id int) (dto.User, error) {
// 	query := `
//         SELECT id, username, email, created_at
//         FROM users
//         WHERE id = $1;
//     `
// 	var u dto.User

// 	err := r.Pool.QueryRow(
// 		query,
// 		id,
// 	).Scan(
// 		&u.ID,
// 		&u.Username,
// 		&u.Email,
// 		&u.CreatedAt,
// 	)

// 	if err != nil {
// 		return dto.User{}, err
// 	}

// 	return u, nil
// }
