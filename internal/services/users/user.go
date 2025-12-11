package users

import (
	"github.com/Craig-Spencer-12/api-demo/internal/entity/dto"
	"github.com/Craig-Spencer-12/api-demo/internal/repo"
)

type Service struct {
	repo repo.UsersRepo
}

func (s *Service) CreateUser(userReq dto.CreateUserRequest) (dto.User, error) {
	return s.repo.Create(userReq)
}

func (s *Service) GetAllUsers() ([]dto.User, error) {
	return s.repo.GetAll()
}

func (s *Service) GetUserByID(id int) (dto.User, error) {
	return s.repo.GetByID(id)
}
