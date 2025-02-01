package task

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	UserId      uuid.UUID
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func NewTask(description, userId string) (*Task, error) {
	userIdAsUuid, err := uuid.Parse(userId)

	if err != nil {
		return nil, err
	}

	return &Task{
		UserId: userIdAsUuid, Description: description,
	}, nil
}
