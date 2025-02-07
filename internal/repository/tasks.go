package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"poymanov/todo/internal/domain"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (repo *TaskRepository) Create(task *domain.Task) (*domain.Task, error) {
	result := repo.db.Create(task)

	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (repo *TaskRepository) Update(task *domain.Task) (*domain.Task, error) {
	result := repo.db.Updates(task)

	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (repo *TaskRepository) Delete(id uuid.UUID) error {
	result := repo.db.Delete(&domain.Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TaskRepository) IsExistsById(id uuid.UUID) bool {
	if err := repo.db.First(&domain.Task{ID: id}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func (repo *TaskRepository) GetAllByUserId(id uuid.UUID) *[]domain.Task {
	var tasks []domain.Task

	repo.db.
		Table("tasks").
		Where("deleted_at is null and user_id = ?", id).
		Order("created_at desc").
		Scan(&tasks)

	return &tasks
}
