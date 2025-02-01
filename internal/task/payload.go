package task

type CreateTaskRequest struct {
	Description string `json:"description" validate:"required"`
}

type UpdateTaskRequest struct {
	Description string `json:"description" validate:"required"`
}
