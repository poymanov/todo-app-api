package service

import (
	"poymanov/todo/internal/domain"
	"poymanov/todo/internal/repository"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(name, email, password string) (*domain.User, error) {
	createdUser, err := s.userRepo.Create(&domain.User{Name: name, Email: email, Password: password})

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) FindByEmail(email string) (*domain.User, error) {
	findUser, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return findUser, nil
}
