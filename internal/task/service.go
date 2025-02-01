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
