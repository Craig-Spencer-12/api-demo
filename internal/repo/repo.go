package repo

import (
	"errors"

	"github.com/Craig-Spencer-12/api-demo/internal/entity/dto"
)

type Repository struct {
	data   []dto.User
	nextID int
}

func NewUsersRepo() *Repository {
	return &Repository{
		data:   []dto.User{},
		nextID: 1,
	}
}

// TODO: Create should return an error not a user but this
// isn't worth changing since I'm switching to a real db at some point
func (r *Repository) Create(userReq dto.CreateUserRequest) dto.User {
	user := dto.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		ID:       r.nextID,
	}

	r.data = append(r.data, user)
	r.nextID++
	return user
}

func (r *Repository) GetAll() []dto.User {
	return r.data
}

func (r *Repository) GetByID(id int) (dto.User, error) {
	for _, user := range r.data {
		if user.ID == id {
			return user, nil
		}
	}
	return dto.User{}, errors.New("message not found")
}
