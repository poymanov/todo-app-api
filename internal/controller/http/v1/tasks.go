package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"poymanov/todo/pkg/response"
	"time"
)

const (
	ErrFailedToGetUser    = "failed to get user"
	ErrFailedToCreateTask = "failed to create task"
	ErrTaskNotFound       = "task not found"
	ErrFailedToUpdateTask = "failed to update task"
	ErrFailedToDeleteTask = "failed to delete task"
)

type CreateTaskRequest struct {
	Description string `json:"description" binding:"required"`
}

type UpdateTaskRequest struct {
	Description string `json:"description" binding:"required"`
}

type GetAllByUserIdResponse struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}

func (h *Handler) initTasksRoutes(api *gin.RouterGroup) {
	tasks := api.Group("/tasks", h.auth)
	{
		tasks.GET("", h.getAllTasksByUserId)
		tasks.POST("", h.createTask)
		tasks.PATCH("/:id", h.updateTaskDescription)
		tasks.PATCH("/:id/complete", h.updateTaskIsComplete(true))
		tasks.PATCH("/:id/incomplete", h.updateTaskIsComplete(false))
		tasks.DELETE("/:id", h.deleteTask)
	}
}

// @Description	Создание задачи
// @Tags			task
// @Param			data	body	CreateTaskRequest	true	"Данные новой задачи"
// @Success		204
// @Failure		400	{object}	response.ErrorResponse
// @Failure		422	{object}	response.ErrorResponse
// @Security		ApiKeyAuth
// @Router			/tasks [post]
func (h *Handler) createTask(c *gin.Context) {
	var body CreateTaskRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userEmail, err := getContextEmail(c)

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToGetUser)
		return
	}

	existedUser, _ := h.services.User.FindByEmail(userEmail)

	if existedUser == nil {
		response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToGetUser)
		return
	}

	_, err = h.services.Task.Create(body.Description, existedUser.ID)

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToCreateTask)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Description	Обновление задачи
// @Tags			task
// @Param			id		path	string				true	"ID задачи"
// @Param			data	body	UpdateTaskRequest	true	"Новые данные для задачи"
// @Success		204
// @Failure		400	{object}	response.ErrorResponse
// @Failure		404	{object}	response.ErrorResponse
// @Failure		422	{object}	response.ErrorResponse
// @Security		ApiKeyAuth
// @Router			/tasks/{id} [patch]
func (h *Handler) updateTaskDescription(c *gin.Context) {
	var body UpdateTaskRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		response.NewErrorResponse(c, http.StatusNotFound, ErrTaskNotFound)
		return
	}

	if !h.services.Task.IsExistsById(id) {
		response.NewErrorResponse(c, http.StatusNotFound, ErrTaskNotFound)
		return
	}

	_, err = h.services.Task.UpdateDescription(id, body.Description)

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToUpdateTask)
		return
	}

	c.Status(http.StatusNoContent)
}

// @Description	Обновление статуса завершения задачи
// @Tags			task
// @Param			id	path	string	true	"ID задачи"
// @Success		204
// @Failure		400	{object}	response.ErrorResponse
// @Failure		404	{object}	response.ErrorResponse
// @Security		ApiKeyAuth
// @Router			/tasks/{id}/complete [patch]
// @Router			/tasks/{id}/incomplete [patch]
func (h *Handler) updateTaskIsComplete(isComplete bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			response.NewErrorResponse(c, http.StatusNotFound, ErrTaskNotFound)
			return
		}

		if !h.services.Task.IsExistsById(id) {
			response.NewErrorResponse(c, http.StatusNotFound, ErrTaskNotFound)
			return
		}

		_, err = h.services.Task.UpdateIsCompleted(id, isComplete)

		if err != nil {
			response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToUpdateTask)
			return
		}

		c.Status(http.StatusNoContent)
	}

}

// @Description	Удаление задачи
// @Tags			task
// @Param			id	path	string	true	"ID задачи"
// @Success		204
// @Failure		400	{object}	response.ErrorResponse
// @Failure		404	{object}	response.ErrorResponse
// @Security		ApiKeyAuth
// @Router			/tasks/{id} [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		response.NewErrorResponse(c, http.StatusNotFound, ErrTaskNotFound)
		return
	}

	if !h.services.Task.IsExistsById(id) {
		response.NewErrorResponse(c, http.StatusNotFound, ErrTaskNotFound)
		return
	}

	err = h.services.Task.Delete(id)

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToDeleteTask)
		return
	}

	c.Status(http.StatusNoContent)
}

// @Description	Получение списка задач пользователя
// @Tags			task
// @Success		200	{array}		GetAllByUserIdResponse
// @Failure		400	{object}	response.ErrorResponse
// @Router			/tasks [get]
func (h *Handler) getAllTasksByUserId(c *gin.Context) {
	userEmail, err := getContextEmail(c)

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToGetUser)
		return
	}

	existedUser, err := h.services.User.FindByEmail(userEmail)

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToGetUser)
		return
	}

	tasks := h.services.Task.GetAllByUserId(existedUser.ID)

	var tasksResponse = make([]GetAllByUserIdResponse, 0)

	for _, task := range *tasks {
		tasksResponse = append(tasksResponse, GetAllByUserIdResponse{
			Id:          task.ID.String(),
			Description: task.Description,
			IsCompleted: *task.IsCompleted,
			CreatedAt:   task.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, tasksResponse)
}
