package task

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"poymanov/todo/internal/service"
	"poymanov/todo/pkg/jwt"
	"poymanov/todo/pkg/middleware"
	"poymanov/todo/pkg/request"
	"poymanov/todo/pkg/response"
)

const (
	ErrFailedToGetUser    = "failed to get user"
	ErrFailedToCreateTask = "failed to create task"
	ErrTaskNotFound       = "task not found"
	ErrFailedToUpdateTask = "failed to update task"
	ErrFailedToDeleteTask = "failed to delete task"
)

type TaskHandler struct {
	services *service.Services
	jwt      *jwt.JWT
}

func NewTaskHandler(router *http.ServeMux, services *service.Services, jwt *jwt.JWT) {
	handler := &TaskHandler{
		services: services,
		jwt:      jwt,
	}
	router.Handle("POST /tasks", middleware.Auth(handler.create(), jwt))
	router.Handle("GET /tasks", middleware.Auth(handler.getAllByUserId(), jwt))
	router.Handle("PATCH /tasks/{id}", middleware.Auth(handler.updateDescription(), jwt))
	router.Handle("PATCH /tasks/{id}/complete", middleware.Auth(handler.updateIsComplete(true), jwt))
	router.Handle("PATCH /tasks/{id}/incomplete", middleware.Auth(handler.updateIsComplete(false), jwt))
	router.Handle("DELETE /tasks/{id}", middleware.Auth(handler.delete(), jwt))
}

// @Description	Создание задачи
// @Tags			task
// @Param			data	body	task.CreateTaskRequest	true	"Данные новой задачи"
// @Success		204
// @Failure		400	{object}	response.ErrorResponse
// @Failure		422	{object}	response.ErrorResponse
// @Security		ApiKeyAuth
// @Router			/tasks [post]
func (h *TaskHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[CreateTaskRequest](req)

		if err != nil {
			response.JsonError(w, err, http.StatusUnprocessableEntity)
			return
		}

		userEmail, ok := req.Context().Value(middleware.ContextEmailKey).(string)

		if !ok {
			response.JsonError(w, errors.New(ErrFailedToGetUser), http.StatusBadRequest)
		}

		existedUser, _ := h.services.User.FindByEmail(userEmail)

		if existedUser == nil {
			response.JsonError(w, errors.New(ErrFailedToGetUser), http.StatusBadRequest)
			return
		}

		_, err = h.services.Task.Create(body.Description, existedUser.ID)

		if err != nil {
			response.JsonError(w, errors.New(ErrFailedToCreateTask), http.StatusBadRequest)
		}

		response.NoContent(w)
	}
}

// @Description	Обновление задачи
// @Tags			task
// @Param			id		path	string					true	"ID задачи"
// @Param			data	body	task.UpdateTaskRequest	true	"Новые данные для задачи"
// @Success		204
// @Failure		400	{object}	response.ErrorResponse
// @Failure		404	{object}	response.ErrorResponse
// @Failure		422	{object}	response.ErrorResponse
// @Security		ApiKeyAuth
// @Router			/tasks/{id} [patch]
func (h *TaskHandler) updateDescription() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		body, err := request.HandleBody[UpdateTaskRequest](req)

		if err != nil {
			response.JsonError(w, err, http.StatusUnprocessableEntity)
			return
		}

		idAsUuid, err := uuid.Parse(id)

		if err != nil {
			response.JsonError(w, errors.New(ErrTaskNotFound), http.StatusNotFound)
			return
		}

		if !h.services.Task.IsExistsById(idAsUuid) {
			response.JsonError(w, errors.New(ErrTaskNotFound), http.StatusNotFound)
			return
		}

		_, err = h.services.Task.UpdateDescription(idAsUuid, body.Description)

		if err != nil {
			response.JsonError(w, errors.New(ErrFailedToUpdateTask), http.StatusBadRequest)
			return
		}

		response.NoContent(w)
	}
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
func (h *TaskHandler) updateIsComplete(isComplete bool) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		idAsUuid, err := uuid.Parse(id)

		if err != nil {
			response.JsonError(w, errors.New(ErrTaskNotFound), http.StatusNotFound)
			return
		}

		if !h.services.Task.IsExistsById(idAsUuid) {
			response.JsonError(w, errors.New(ErrTaskNotFound), http.StatusNotFound)
			return
		}

		_, err = h.services.Task.UpdateIsCompleted(idAsUuid, isComplete)

		if err != nil {
			response.JsonError(w, errors.New(ErrFailedToUpdateTask), http.StatusBadRequest)
			return
		}

		response.NoContent(w)
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
func (h *TaskHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		idAsUuid, err := uuid.Parse(id)

		if err != nil {
			response.JsonError(w, errors.New(ErrTaskNotFound), http.StatusNotFound)
			return
		}

		if !h.services.Task.IsExistsById(idAsUuid) {
			response.JsonError(w, errors.New(ErrTaskNotFound), http.StatusNotFound)
			return
		}

		err = h.services.Task.Delete(idAsUuid)

		if err != nil {
			response.JsonError(w, errors.New(ErrFailedToDeleteTask), http.StatusBadRequest)
			return
		}

		response.NoContent(w)
	}
}

// @Description	Получение списка задач пользователя
// @Tags			task
// @Success		200	{array}		task.GetAllByUserIdResponse
// @Failure		400	{object}	response.ErrorResponse
// @Router			/tasks [get]
func (h *TaskHandler) getAllByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		userEmail, ok := req.Context().Value(middleware.ContextEmailKey).(string)

		if !ok {
			response.JsonError(w, errors.New(ErrFailedToGetUser), http.StatusBadRequest)
		}

		existedUser, err := h.services.User.FindByEmail(userEmail)

		if err != nil {
			response.JsonError(w, errors.New(ErrFailedToGetUser), http.StatusBadRequest)
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

		response.Json(w, tasksResponse, http.StatusOK)
	}
}
