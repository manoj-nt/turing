// user_service.go
package service

import "app/domain"

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser(user domain.User) error {
	return s.repo.Save(user)
}

func (s *UserService) ListUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}
