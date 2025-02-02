package task

import (
	"errors"
	"github.com/google/uuid"
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

func (repo *TaskRepository) Update(task *db.Task) (*db.Task, error) {
	result := repo.Db.Updates(task)

	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (repo *TaskRepository) Delete(id uuid.UUID) error {
	result := repo.Db.Delete(&db.Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TaskRepository) IsExistsById(id uuid.UUID) bool {
	if err := repo.Db.First(&db.Task{ID: id}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func (repo *TaskRepository) GetAllByUserId(id uuid.UUID) *[]db.Task {
	var tasks []db.Task

	repo.Db.
		Table("tasks").
		Where("deleted_at is null and user_id = ?", id).
		Order("created_at desc").
		Scan(&tasks)

	return &tasks
}
