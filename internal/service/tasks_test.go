package service_test

import (
	"errors"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mock_repository "poymanov/todo/internal/repository/mocks"
	"poymanov/todo/internal/service"
	"poymanov/todo/pkg/db"
	"testing"
)

func TestTaskServiceCreate_Failed(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskRepo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("failed"))

	userId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	createdTask, err := taskService.Create(faker.Word(), userId)

	require.Error(t, err)
	require.Nil(t, createdTask)
}

func TestTaskServiceCreate_Success(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	userId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskData := db.Task{Description: faker.Word(), UserId: userId}

	taskRepo.EXPECT().Create(gomock.Any()).Return(&taskData, nil)

	createdTask, err := taskService.Create(faker.Word(), userId)

	require.NoError(t, err)
	require.NotNil(t, createdTask)
	require.Equal(t, taskData.Description, createdTask.Description)
	require.Equal(t, userId, createdTask.UserId)
}

func TestTaskServiceUpdateDescription_Failed(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepo.EXPECT().Update(gomock.Any()).Return(nil, errors.New("failed"))

	updatedTask, err := taskService.UpdateDescription(taskId, faker.Word())

	require.Error(t, err)
	require.Nil(t, updatedTask)
}

func TestTaskServiceUpdateDescription_Success(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	newDescription := faker.Word()
	taskData := db.Task{ID: taskId, Description: newDescription}

	taskRepo.EXPECT().Update(gomock.Any()).Return(&taskData, nil)

	updatedTask, err := taskService.UpdateDescription(taskId, newDescription)

	require.NoError(t, err)
	require.NotNil(t, updatedTask)
	require.Equal(t, newDescription, updatedTask.Description)
}

func TestTaskServiceUpdateIsCompleted_Failed(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepo.EXPECT().Update(gomock.Any()).Return(nil, errors.New("failed"))

	updatedTask, err := taskService.UpdateIsCompleted(taskId, false)

	require.Error(t, err)
	require.Nil(t, updatedTask)
}

func TestTaskServiceUpdateIsCompleted_Completed(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	isCompleted := true

	taskData := db.Task{ID: taskId, IsCompleted: &isCompleted}

	taskRepo.EXPECT().Update(gomock.Any()).Return(&taskData, nil)

	updatedTask, err := taskService.UpdateIsCompleted(taskId, true)

	require.NoError(t, err)
	require.NotNil(t, updatedTask)
	require.True(t, *updatedTask.IsCompleted)
}

func TestTaskServiceUpdateIsCompleted_NotCompleted(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	isCompleted := false

	taskData := db.Task{ID: taskId, IsCompleted: &isCompleted}

	taskRepo.EXPECT().Update(gomock.Any()).Return(&taskData, nil)

	updatedTask, err := taskService.UpdateIsCompleted(taskId, true)

	require.NoError(t, err)
	require.NotNil(t, updatedTask)
	require.False(t, *updatedTask.IsCompleted)
}

func TestTaskServiceDelete_Failed(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepo.EXPECT().Delete(gomock.Any()).Return(errors.New("failed"))

	err = taskService.Delete(taskId)

	require.Error(t, err)
}

func TestTaskServiceDelete_Success(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepo.EXPECT().Delete(gomock.Any()).Return(nil)

	err = taskService.Delete(taskId)

	require.NoError(t, err)
}

func TestTaskServiceIsExistsById_NotExists(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepo.EXPECT().IsExistsById(gomock.Any()).Return(false)

	isExists := taskService.IsExistsById(taskId)

	require.False(t, isExists)
}

func TestTaskServiceIsExistsById_Exists(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepo.EXPECT().IsExistsById(gomock.Any()).Return(true)

	isExists := taskService.IsExistsById(taskId)

	require.True(t, isExists)
}

func TestTaskServiceGetAllByUserId_Empty(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	userId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepo.EXPECT().GetAllByUserId(gomock.Any()).Return(&[]db.Task{})

	tasks := taskService.GetAllByUserId(userId)

	require.Empty(t, tasks)
}

func TestTaskServiceGetAllByUserId_Success(t *testing.T) {
	taskService, taskRepo := mockTaskService(t)

	userId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepo.EXPECT().GetAllByUserId(gomock.Any()).Return(&[]db.Task{{}})

	tasks := taskService.GetAllByUserId(userId)

	require.NotEmpty(t, tasks)
}

func mockTaskService(t *testing.T) (*service.TaskService, *mock_repository.MockTask) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	taskRepo := mock_repository.NewMockTask(mockCtl)

	taskService := service.NewTaskService(taskRepo)

	return taskService, taskRepo
}
