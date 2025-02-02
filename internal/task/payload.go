package task

import "time"

type CreateTaskRequest struct {
	Description string `json:"description" validate:"required"`
}

type UpdateTaskRequest struct {
	Description string `json:"description" validate:"required"`
}

type GetAllByUserIdResponse struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}
