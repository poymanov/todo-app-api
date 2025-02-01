package task

import (
	"gorm.io/gorm"
	"poymanov/todo/pkg/db"
)

type TaskRepository struct {
	Db *gorm.DB
}

type TaskRepositoryDeps struct {
	Db *gorm.DB
}

func NewTaskRepository(deps TaskRepositoryDeps) *TaskRepository {
	return &TaskRepository{Db: deps.Db}
}

func (repo *TaskRepository) Create(task *db.Task) (*db.Task, error) {
	result := repo.Db.Create(task)

	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}
