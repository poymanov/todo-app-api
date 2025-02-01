package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"poymanov/todo/internal/task"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Tasks     []task.Task    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
