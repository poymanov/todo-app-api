package service

import (
	"github.com/google/uuid"
	"poymanov/todo/internal/domain"
	"poymanov/todo/internal/repository"
)

type TaskService struct {
	taskRepo repository.Task
}

func NewTaskService(taskRepo repository.Task) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) Create(description string, userId uuid.UUID) (*domain.Task, error) {
	createdTask, err := s.taskRepo.Create(&domain.Task{Description: description, UserId: userId})

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (s *TaskService) UpdateDescription(id uuid.UUID, description string) (*domain.Task, error) {
	updatedTask, err := s.taskRepo.Update(&domain.Task{
		ID: id, Description: description,
	})

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (s *TaskService) UpdateIsCompleted(id uuid.UUID, isCompleted bool) (*domain.Task, error) {
	updatedTask, err := s.taskRepo.Update(&domain.Task{
		ID: id, IsCompleted: &isCompleted,
	})

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (s *TaskService) Delete(id uuid.UUID) error {
	result := s.taskRepo.Delete(id)

	if result != nil {
		return result
	}

	return nil
}

func (s *TaskService) IsExistsById(id uuid.UUID) bool {
	return s.taskRepo.IsExistsById(id)
}

func (s *TaskService) GetAllByUserId(id uuid.UUID) *[]domain.Task {
	return s.taskRepo.GetAllByUserId(id)
}
