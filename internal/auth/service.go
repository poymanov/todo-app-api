package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"poymanov/todo/internal/user"
	"poymanov/todo/pkg/jwt"
)

type AuthService struct {
	UserService *user.UserService
	JWT         *jwt.JWT
}

type AuthServiceDeps struct {
	UserService *user.UserService
	JWT         *jwt.JWT
}

func NewAuthService(deps AuthServiceDeps) *AuthService {
	return &AuthService{UserService: deps.UserService, JWT: deps.JWT}
}

func (s *AuthService) Register(data RegisterData) (string, error) {
	existedUser, _ := s.UserService.FindByEmail(data.Email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	createdUser, err := s.UserService.Create(data.Name, data.Email, string(hashedPassword))

	if err != nil {
		return "", err
	}

	token, err := s.JWT.Create(jwt.JWTData{
		Email: createdUser.Email,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}
