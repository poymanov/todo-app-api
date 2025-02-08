package v1

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"poymanov/todo/internal/domain"
	"poymanov/todo/internal/service"
	mock_service "poymanov/todo/internal/service/mocks"
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {
	testCases := []struct {
		name            string
		body            string
		response        string
		statusCode      int
		contextModifier func(c *gin.Context)
		mockFunction    func(userService *mock_service.MockUser, taskService *mock_service.MockTask)
	}{
		{
			name:            "Empty",
			body:            ``,
			response:        `{"message":"EOF"}`,
			statusCode:      http.StatusUnprocessableEntity,
			contextModifier: func(c *gin.Context) {},
			mockFunction:    func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {},
		},
		{
			name:            "Missing description",
			body:            `{}`,
			response:        `{"message":"Key: 'CreateTaskRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag"}`,
			statusCode:      http.StatusUnprocessableEntity,
			contextModifier: func(c *gin.Context) {},
			mockFunction:    func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {},
		},
		{
			name:            "Failed to get email from context",
			body:            `{"description": "test"}`,
			response:        `{"message":"Failed to get user"}`,
			statusCode:      http.StatusBadRequest,
			contextModifier: func(c *gin.Context) {},
			mockFunction:    func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {},
		},
		{
			name:       "Not existed user",
			body:       `{"description": "test"}`,
			response:   `{"message":"Failed to get user"}`,
			statusCode: http.StatusBadRequest,
			contextModifier: func(c *gin.Context) {
				c.Set(ContextEmailKey, faker.Email())
			},
			mockFunction: func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {
				userService.EXPECT().FindByEmail(gomock.Any()).Return(nil, errors.New("failed"))
			},
		},
		{
			name:       "Failed to create task",
			body:       `{"description": "test"}`,
			response:   `{"message":"Failed to create task"}`,
			statusCode: http.StatusBadRequest,
			contextModifier: func(c *gin.Context) {
				c.Set(ContextEmailKey, faker.Email())
			},
			mockFunction: func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {
				userService.EXPECT().FindByEmail(gomock.Any()).Return(&domain.User{}, nil)
				taskService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("failed"))
			},
		},
		{
			name:       "Success",
			body:       `{"description": "test"}`,
			response:   ``,
			statusCode: http.StatusNoContent,
			contextModifier: func(c *gin.Context) {
				c.Set(ContextEmailKey, faker.Email())
			},
			mockFunction: func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {
				userService.EXPECT().FindByEmail(gomock.Any()).Return(&domain.User{}, nil)
				taskService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&domain.Task{}, nil)
			},
		},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userService := mock_service.NewMockUser(c)
			taskService := mock_service.NewMockTask(c)

			tc.mockFunction(userService, taskService)
			handler := Handler{services: &service.Services{User: userService, Task: taskService}}

			r := gin.New()
			r.POST("/tasks", tc.contextModifier, handler.createTask)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/tasks", bytes.NewBufferString(tc.body))
			r.ServeHTTP(w, req)

			require.Equal(t, tc.statusCode, w.Code)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}

func TestUpdateTaskDescription(t *testing.T) {
	testCases := []struct {
		name         string
		body         string
		taskId       string
		response     string
		statusCode   int
		mockFunction func(taskService *mock_service.MockTask)
	}{
		{
			name:         "Empty",
			body:         ``,
			taskId:       faker.UUIDHyphenated(),
			response:     `{"message":"EOF"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(taskService *mock_service.MockTask) {},
		},
		{
			name:         "Missing description",
			body:         `{}`,
			taskId:       faker.UUIDHyphenated(),
			response:     `{"message":"Key: 'UpdateTaskRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(taskService *mock_service.MockTask) {},
		},
		{
			name:         "Failed to parse task id",
			body:         `{"description": "test"}`,
			taskId:       faker.Word(),
			response:     `{"message":"Task not found"}`,
			statusCode:   http.StatusNotFound,
			mockFunction: func(taskService *mock_service.MockTask) {},
		},
		{
			name:       "Task not existed",
			body:       `{"description": "test"}`,
			taskId:     faker.UUIDHyphenated(),
			response:   `{"message":"Task not found"}`,
			statusCode: http.StatusNotFound,
			mockFunction: func(taskService *mock_service.MockTask) {
				taskService.EXPECT().IsExistsById(gomock.Any()).Return(false)
			},
		},
		{
			name:       "Success",
			body:       `{"description": "test"}`,
			taskId:     faker.UUIDHyphenated(),
			response:   ``,
			statusCode: http.StatusNoContent,
			mockFunction: func(taskService *mock_service.MockTask) {
				taskService.EXPECT().IsExistsById(gomock.Any()).Return(true)
				taskService.EXPECT().UpdateDescription(gomock.Any(), gomock.Any()).Return(&domain.Task{}, nil)
			},
		},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			taskService := mock_service.NewMockTask(c)

			tc.mockFunction(taskService)
			handler := Handler{services: &service.Services{Task: taskService}}

			r := gin.New()
			r.PATCH("/tasks/:id", handler.updateTaskDescription)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", "/tasks/"+tc.taskId, bytes.NewBufferString(tc.body))
			r.ServeHTTP(w, req)

			require.Equal(t, tc.statusCode, w.Code)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}

func TestUpdateTaskIsComplete(t *testing.T) {
	testCases := []struct {
		name         string
		taskId       string
		response     string
		statusCode   int
		isComplete   bool
		routerPath   string
		requestPath  string
		mockFunction func(taskService *mock_service.MockTask)
	}{
		{
			name:         "Failed to parse task id (incomplete)",
			taskId:       faker.Word(),
			response:     `{"message":"Task not found"}`,
			statusCode:   http.StatusNotFound,
			isComplete:   false,
			routerPath:   "/tasks/:id/incomplete",
			requestPath:  fmt.Sprintf("/tasks/%s/incomplete", faker.Word()),
			mockFunction: func(taskService *mock_service.MockTask) {},
		},
		{
			name:        "Task not existed (incomplete)",
			taskId:      faker.UUIDHyphenated(),
			response:    `{"message":"Task not found"}`,
			statusCode:  http.StatusNotFound,
			isComplete:  false,
			routerPath:  "/tasks/:id/incomplete",
			requestPath: fmt.Sprintf("/tasks/%s/incomplete", faker.UUIDHyphenated()),
			mockFunction: func(taskService *mock_service.MockTask) {
				taskService.EXPECT().IsExistsById(gomock.Any()).Return(false)
			},
		},
		{
			name:        "Success (incomplete)",
			taskId:      faker.UUIDHyphenated(),
			response:    ``,
			statusCode:  http.StatusNoContent,
			isComplete:  false,
			routerPath:  "/tasks/:id/incomplete",
			requestPath: fmt.Sprintf("/tasks/%s/incomplete", faker.UUIDHyphenated()),
			mockFunction: func(taskService *mock_service.MockTask) {
				taskService.EXPECT().IsExistsById(gomock.Any()).Return(true)
				taskService.EXPECT().UpdateIsCompleted(gomock.Any(), gomock.Any()).Return(&domain.Task{}, nil)
			},
		},
		{
			name:         "Failed to parse task id (complete)",
			taskId:       faker.Word(),
			response:     `{"message":"Task not found"}`,
			statusCode:   http.StatusNotFound,
			isComplete:   false,
			routerPath:   "/tasks/:id/complete",
			requestPath:  fmt.Sprintf("/tasks/%s/complete", faker.Word()),
			mockFunction: func(taskService *mock_service.MockTask) {},
		},
		{
			name:        "Task not existed (complete)",
			taskId:      faker.UUIDHyphenated(),
			response:    `{"message":"Task not found"}`,
			statusCode:  http.StatusNotFound,
			isComplete:  false,
			routerPath:  "/tasks/:id/complete",
			requestPath: fmt.Sprintf("/tasks/%s/complete", faker.UUIDHyphenated()),
			mockFunction: func(taskService *mock_service.MockTask) {
				taskService.EXPECT().IsExistsById(gomock.Any()).Return(false)
			},
		},
		{
			name:        "Success (complete)",
			taskId:      faker.UUIDHyphenated(),
			response:    ``,
			statusCode:  http.StatusNoContent,
			isComplete:  false,
			routerPath:  "/tasks/:id/complete",
			requestPath: fmt.Sprintf("/tasks/%s/complete", faker.UUIDHyphenated()),
			mockFunction: func(taskService *mock_service.MockTask) {
				taskService.EXPECT().IsExistsById(gomock.Any()).Return(true)
				taskService.EXPECT().UpdateIsCompleted(gomock.Any(), gomock.Any()).Return(&domain.Task{}, nil)
			},
		},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			taskService := mock_service.NewMockTask(c)

			tc.mockFunction(taskService)
			handler := Handler{services: &service.Services{Task: taskService}}

			r := gin.New()
			r.PATCH(tc.routerPath, handler.updateTaskIsComplete(tc.isComplete))

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PATCH", tc.requestPath, nil)
			r.ServeHTTP(w, req)

			require.Equal(t, tc.statusCode, w.Code)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}

func TestDeleteTask(t *testing.T) {
	testCases := []struct {
		name         string
		taskId       string
		response     string
		statusCode   int
		mockFunction func(taskService *mock_service.MockTask)
	}{
		{
			name:         "Failed to parse task id",
			taskId:       faker.Word(),
			response:     `{"message":"Task not found"}`,
			statusCode:   http.StatusNotFound,
			mockFunction: func(taskService *mock_service.MockTask) {},
		},
		{
			name:       "Task not existed",
			taskId:     faker.UUIDHyphenated(),
			response:   `{"message":"Task not found"}`,
			statusCode: http.StatusNotFound,
			mockFunction: func(taskService *mock_service.MockTask) {
				taskService.EXPECT().IsExistsById(gomock.Any()).Return(false)
			},
		},
		{
			name:       "Failed to delete task",
			taskId:     faker.UUIDHyphenated(),
			response:   `{"message":"Failed to delete task"}`,
			statusCode: http.StatusBadRequest,
			mockFunction: func(taskService *mock_service.MockTask) {
				taskService.EXPECT().IsExistsById(gomock.Any()).Return(true)
				taskService.EXPECT().Delete(gomock.Any()).Return(errors.New("failed"))
			},
		},
		{
			name:       "Success",
			taskId:     faker.UUIDHyphenated(),
			response:   ``,
			statusCode: http.StatusNoContent,
			mockFunction: func(taskService *mock_service.MockTask) {
				taskService.EXPECT().IsExistsById(gomock.Any()).Return(true)
				taskService.EXPECT().Delete(gomock.Any()).Return(nil)
			},
		},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			taskService := mock_service.NewMockTask(c)

			tc.mockFunction(taskService)
			handler := Handler{services: &service.Services{Task: taskService}}

			r := gin.New()
			r.DELETE("/tasks/:id", handler.deleteTask)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/tasks/"+tc.taskId, nil)
			r.ServeHTTP(w, req)

			require.Equal(t, tc.statusCode, w.Code)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}

func TestGetAllTasksByUserId(t *testing.T) {
	testCases := []struct {
		name            string
		response        string
		statusCode      int
		contextModifier func(c *gin.Context)
		mockFunction    func(userService *mock_service.MockUser, taskService *mock_service.MockTask)
	}{
		{
			name:            "Failed to get email from context",
			response:        `{"message":"Failed to get user"}`,
			statusCode:      http.StatusBadRequest,
			contextModifier: func(c *gin.Context) {},
			mockFunction:    func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {},
		},
		{
			name:       "Not existed user",
			response:   `{"message":"Failed to get user"}`,
			statusCode: http.StatusBadRequest,
			contextModifier: func(c *gin.Context) {
				c.Set(ContextEmailKey, faker.Email())
			},
			mockFunction: func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {
				userService.EXPECT().FindByEmail(gomock.Any()).Return(nil, errors.New("failed"))
			},
		},
		{
			name:       "Tasks no exists",
			response:   `[]`,
			statusCode: http.StatusOK,
			contextModifier: func(c *gin.Context) {
				c.Set(ContextEmailKey, faker.Email())
			},
			mockFunction: func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {
				userService.EXPECT().FindByEmail(gomock.Any()).Return(&domain.User{}, nil)
				taskService.EXPECT().GetAllByUserId(gomock.Any()).Return(&[]domain.Task{})
			},
		},
		{
			name:       "Success",
			response:   `[{"id":"8d306d55-4301-4770-8a90-e64f771dc3f9","description":"Description","is_completed":true,"created_at":"2006-01-02T15:04:05Z"}]`,
			statusCode: http.StatusOK,
			contextModifier: func(c *gin.Context) {
				c.Set(ContextEmailKey, faker.Email())
			},
			mockFunction: func(userService *mock_service.MockUser, taskService *mock_service.MockTask) {
				taskId, _ := uuid.Parse("8d306d55-4301-4770-8a90-e64f771dc3f9")
				isCompleted := true
				createdAt, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")

				userService.EXPECT().FindByEmail(gomock.Any()).Return(&domain.User{}, nil)
				taskService.EXPECT().GetAllByUserId(gomock.Any()).Return(&[]domain.Task{
					{
						ID:          taskId,
						Description: "Description",
						IsCompleted: &isCompleted,
						CreatedAt:   createdAt,
					},
				})
			},
		},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userService := mock_service.NewMockUser(c)
			taskService := mock_service.NewMockTask(c)

			tc.mockFunction(userService, taskService)
			handler := Handler{services: &service.Services{User: userService, Task: taskService}}

			r := gin.New()
			r.GET("/tasks", tc.contextModifier, handler.getAllTasksByUserId)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/tasks", nil)
			r.ServeHTTP(w, req)

			require.Equal(t, tc.statusCode, w.Code)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}
