package user

import "poymanov/todo/pkg/db"

type UserService struct {
	UserRepository *UserRepository
}

type UserServiceDeps struct {
	UserRepository *UserRepository
}

func NewUserService(deps UserServiceDeps) *UserService {
	return &UserService{UserRepository: deps.UserRepository}
}

func (s *UserService) Create(name, email, password string) (*db.User, error) {
	newUser := db.NewUser(name, email, password)

	createdUser, err := s.UserRepository.Create(newUser)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) FindByEmail(email string) (*db.User, error) {
	user, err := s.UserRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
