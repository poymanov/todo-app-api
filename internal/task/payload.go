package task

type CreateTaskRequest struct {
	Description string `json:"description" validate:"required"`
}
