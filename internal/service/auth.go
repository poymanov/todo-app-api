package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"poymanov/todo/pkg/jwt"
)

const (
	ErrUserExists       = "user exists"
	ErrWrongCredentials = "wrong email or password"
)

type RegisterData struct {
	Name     string
	Email    string
	Password string
}

type LoginData struct {
	Email    string
	Password string
}

type AuthService struct {
	UserService *UserService
	JWT         *jwt.JWT
}

func NewAuthService(UserService *UserService, JWT *jwt.JWT) *AuthService {
	return &AuthService{UserService: UserService, JWT: JWT}
}

func (s *AuthService) Register(data RegisterData) (string, error) {
	existedUser, _ := s.UserService.FindByEmail(data.Email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

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

func (s *AuthService) Login(data LoginData) (string, error) {
	existedUser, _ := s.UserService.FindByEmail(data.Email)

	if existedUser == nil {
		return "", errors.New(ErrWrongCredentials)
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(data.Password))

	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}

	token, err := s.JWT.Create(jwt.JWTData{
		Email: existedUser.Email,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}
