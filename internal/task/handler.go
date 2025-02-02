package task

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"poymanov/todo/internal/user"
	"poymanov/todo/pkg/jwt"
	"poymanov/todo/pkg/middleware"
	"poymanov/todo/pkg/request"
	"poymanov/todo/pkg/response"
)

type TaskHandlerDeps struct {
	TaskService *TaskService
	UserService *user.UserService
	JWT         *jwt.JWT
}

type TaskHandler struct {
	TaskService *TaskService
	UserService *user.UserService
}

func NewTaskHandler(router *http.ServeMux, deps TaskHandlerDeps) {
	handler := &TaskHandler{
		TaskService: deps.TaskService,
		UserService: deps.UserService,
	}
	router.Handle("POST /tasks", middleware.Auth(handler.create(), deps.JWT))
	router.Handle("GET /tasks", middleware.Auth(handler.getAllByUserId(), deps.JWT))
	router.Handle("PATCH /tasks/{id}", middleware.Auth(handler.update(), deps.JWT))
	router.Handle("DELETE /tasks/{id}", middleware.Auth(handler.delete(), deps.JWT))
}

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

		existedUser, _ := h.UserService.FindByEmail(userEmail)

		if existedUser == nil {
			response.JsonError(w, errors.New(ErrFailedToGetUser), http.StatusBadRequest)
		}

		_, err = h.TaskService.Create(body.Description, existedUser.ID)

		if err != nil {
			response.JsonError(w, errors.New(ErrFailedToCreateTask), http.StatusBadRequest)
		}

		response.NoContent(w)
	}
}

func (h *TaskHandler) update() http.HandlerFunc {
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

		if !h.TaskService.IsExistsById(idAsUuid) {
			response.JsonError(w, errors.New(ErrTaskNotFound), http.StatusNotFound)
			return
		}

		_, err = h.TaskService.Update(idAsUuid, body.Description)

		if err != nil {
			response.JsonError(w, errors.New(ErrFailedToUpdateTask), http.StatusBadRequest)
			return
		}

		response.NoContent(w)
	}
}

func (h *TaskHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")

		idAsUuid, err := uuid.Parse(id)

		if err != nil {
			response.JsonError(w, errors.New(ErrTaskNotFound), http.StatusNotFound)
			return
		}

		if !h.TaskService.IsExistsById(idAsUuid) {
			response.JsonError(w, errors.New(ErrTaskNotFound), http.StatusNotFound)
			return
		}

		err = h.TaskService.Delete(idAsUuid)

		if err != nil {
			response.JsonError(w, errors.New(ErrFailedToDeleteTask), http.StatusBadRequest)
			return
		}

		response.NoContent(w)
	}
}

func (h *TaskHandler) getAllByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		userEmail, ok := req.Context().Value(middleware.ContextEmailKey).(string)

		if !ok {
			response.JsonError(w, errors.New(ErrFailedToGetUser), http.StatusBadRequest)
		}

		existedUser, err := h.UserService.FindByEmail(userEmail)

		if err != nil {
			response.JsonError(w, errors.New(ErrFailedToGetUser), http.StatusBadRequest)
			return
		}

		tasks := h.TaskService.GetAllByUserId(existedUser.ID)

		var tasksResponse = make([]GetAllByUserIdResponse, 0)

		for _, task := range *tasks {
			tasksResponse = append(tasksResponse, GetAllByUserIdResponse{
				Id:          task.ID.String(),
				Description: task.Description,
				CreatedAt:   task.CreatedAt,
			})
		}

		response.Json(w, tasksResponse, http.StatusOK)
	}
}
