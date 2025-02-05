package service

import (
	"github.com/google/uuid"
	"poymanov/todo/internal/repository"
	"poymanov/todo/pkg/db"
)

type TaskService struct {
	taskRepo repository.Task
}

func NewTaskService(taskRepo repository.Task) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) Create(description string, userId uuid.UUID) (*db.Task, error) {
	newTask := db.NewTask(description, userId)

	createdTask, err := s.taskRepo.Create(newTask)

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (s *TaskService) UpdateDescription(id uuid.UUID, description string) (*db.Task, error) {
	updatedTask, err := s.taskRepo.Update(&db.Task{
		ID: id, Description: description,
	})

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (s *TaskService) UpdateIsCompleted(id uuid.UUID, isCompleted bool) (*db.Task, error) {
	updatedTask, err := s.taskRepo.Update(&db.Task{
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

func (s *TaskService) GetAllByUserId(id uuid.UUID) *[]db.Task {
	return s.taskRepo.GetAllByUserId(id)
}
