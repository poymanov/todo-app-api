package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"poymanov/todo/pkg/db"
)

type Task interface {
	Create(task *db.Task) (*db.Task, error)
	Update(task *db.Task) (*db.Task, error)
	Delete(id uuid.UUID) error
	IsExistsById(id uuid.UUID) bool
	GetAllByUserId(id uuid.UUID) *[]db.Task
}

type User interface {
	Create(user *db.User) (*db.User, error)
	FindByEmail(email string) (*db.User, error)
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
