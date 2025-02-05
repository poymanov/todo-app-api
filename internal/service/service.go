package service

import (
	"github.com/google/uuid"
	"poymanov/todo/internal/repository"
	"poymanov/todo/pkg/db"
	"poymanov/todo/pkg/jwt"
)

type Auth interface {
	Register(data RegisterData) (string, error)
	Login(data LoginData) (string, error)
}

type Task interface {
	Create(description string, userId uuid.UUID) (*db.Task, error)
	UpdateDescription(id uuid.UUID, description string) (*db.Task, error)
	UpdateIsCompleted(id uuid.UUID, isCompleted bool) (*db.Task, error)
	Delete(id uuid.UUID) error
	IsExistsById(id uuid.UUID) bool
	GetAllByUserId(id uuid.UUID) *[]db.Task
}

type User interface {
	Create(name, email, password string) (*db.User, error)
	FindByEmail(email string) (*db.User, error)
}

type Services struct {
	Auth Auth
	Task Task
	User User
}

func NewServices(repos *repository.Repositories, jwt *jwt.JWT) *Services {
	usersService := NewUserService(repos.User)
	authService := NewAuthService(usersService, jwt)
	tasksService := NewTaskService(repos.Task)

	return &Services{
		Auth: authService,
		Task: tasksService,
		User: usersService,
	}
}
