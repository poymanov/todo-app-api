package task

import (
	"github.com/google/uuid"
	"poymanov/todo/pkg/db"
)

type TaskService struct {
	TaskRepository *TaskRepository
}

type TaskServiceDeps struct {
	TaskRepository *TaskRepository
}

func NewTaskService(deps TaskServiceDeps) *TaskService {
	return &TaskService{TaskRepository: deps.TaskRepository}
}

func (s *TaskService) Create(description string, userId uuid.UUID) (*db.Task, error) {
	newTask := db.NewTask(description, userId)

	createdTask, err := s.TaskRepository.Create(newTask)

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (s *TaskService) Update(id uuid.UUID, description string) (*db.Task, error) {
	updatedTask, err := s.TaskRepository.Update(&db.Task{
		ID: id, Description: description,
	})

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (s *TaskService) Delete(id uuid.UUID) error {
	result := s.TaskRepository.Delete(id)

	if result != nil {
		return result
	}

	return nil
}

func (s *TaskService) IsExistsById(id uuid.UUID) bool {
	return s.TaskRepository.IsExistsById(id)
}
