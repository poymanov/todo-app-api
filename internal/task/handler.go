package task

import (
	"errors"
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
