package repo

import (
	"github.com/Craig-Spencer-12/api-demo/internal/entity/dto"
	"github.com/Craig-Spencer-12/api-demo/pkg/db"
)

type UsersRepo struct {
	*db.SQL
}

func NewUsersRepo(database *db.SQL) *UsersRepo {
	return &UsersRepo{database}
}

func (r *UsersRepo) Create(userReq dto.CreateUserRequest) (dto.User, error) {
	var user dto.User

	query := `
        INSERT INTO users (username, email)
        VALUES ($1, $2)
        RETURNING id, username, email, created_at;
    `

	err := r.Pool.QueryRow(
		query,
		userReq.Username,
		userReq.Email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		return dto.User{}, err
	}

	return user, nil
}

func (r *UsersRepo) GetAll() ([]dto.User, error) {
	query := `
        SELECT id, username, email, created_at
        FROM users
        ORDER BY id ASC;
    `
	rows, err := r.Pool.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []dto.User{}

	for rows.Next() {
		var u dto.User
		err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *UsersRepo) GetByID(id int) (dto.User, error) {
	query := `
        SELECT id, username, email, created_at
        FROM users
        WHERE id = $1;
    `
	var u dto.User

	err := r.Pool.QueryRow(
		query,
		id,
	).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.CreatedAt,
	)

	if err != nil {
		return dto.User{}, err
	}

	return u, nil
}
