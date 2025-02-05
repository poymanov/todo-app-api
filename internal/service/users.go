package service

import (
	"poymanov/todo/internal/repository"
	"poymanov/todo/pkg/db"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(name, email, password string) (*db.User, error) {
	newUser := db.NewUser(name, email, password)

	createdUser, err := s.userRepo.Create(newUser)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) FindByEmail(email string) (*db.User, error) {
	findUser, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return findUser, nil
}
