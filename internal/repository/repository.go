package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"poymanov/todo/internal/domain"
)

type Task interface {
	Create(task *domain.Task) (*domain.Task, error)
	Update(task *domain.Task) (*domain.Task, error)
	Delete(id uuid.UUID) error
	IsExistsById(id uuid.UUID) bool
	GetAllByUserId(id uuid.UUID) *[]domain.Task
}

type User interface {
	Create(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

type Repositories struct {
	Task Task
	User User
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Task: NewTaskRepository(db),
		User: NewUserRepository(db),
	}
}
