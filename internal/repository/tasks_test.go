package repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"poymanov/todo/internal/repository"
	"poymanov/todo/pkg/db"
	"poymanov/todo/pkg/helpers"
	"testing"
)

func TestTaskRepositoryCreateSuccess(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	taskUuid := faker.UUIDHyphenated()
	userUuid, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(taskUuid))
	mock.ExpectCommit()

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	expectedTask := db.Task{UserId: userUuid, Description: faker.Word()}

	createdTask, err := taskRepository.Create(&expectedTask)

	require.NoError(t, err)

	require.Equal(t, expectedTask.ID.String(), taskUuid)
	require.Equal(t, expectedTask.Description, createdTask.Description)
	require.Equal(t, expectedTask.UserId.String(), createdTask.UserId.String())
	require.False(t, *createdTask.IsCompleted)
}

func TestTaskRepositoryCreateFailed(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnError(gorm.ErrInvalidValue)
	mock.ExpectRollback()

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	newUser := db.Task{Description: faker.Word()}

	createdUser, err := taskRepository.Create(&newUser)

	require.Nil(t, createdUser)
	require.Error(t, err)
	require.Equal(t, gorm.ErrInvalidValue, err)
}

func TestTaskRepositoryUpdateSuccess(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	newDescription := faker.Word()

	require.NoError(t, err)

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	taskUpdate, err := taskRepository.Update(&db.Task{ID: taskId, Description: newDescription})

	require.NoError(t, err)
	require.Equal(t, taskUpdate.ID, taskId)
	require.Equal(t, taskUpdate.Description, newDescription)
}

func TestTaskRepositoryUpdateFailed(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	taskId, err := uuid.Parse(faker.UUIDHyphenated())

	require.NoError(t, err)

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnError(gorm.ErrInvalidValue)
	mock.ExpectRollback()

	taskUpdate, err := taskRepository.Update(&db.Task{ID: taskId, Description: faker.Word()})

	require.Nil(t, taskUpdate)
	require.Error(t, err)
	require.Equal(t, gorm.ErrInvalidValue, err)
}

func TestTaskRepositoryDeleteSuccess(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	taskId, err := uuid.Parse(faker.UUIDHyphenated())

	require.NoError(t, err)

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err = taskRepository.Delete(taskId)

	require.NoError(t, err)
}

func TestTaskRepositoryDeleteFailed(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	taskId, err := uuid.Parse(faker.UUIDHyphenated())

	require.NoError(t, err)

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnError(gorm.ErrInvalidValue)
	mock.ExpectRollback()

	err = taskRepository.Delete(taskId)

	require.Error(t, err)
	require.Equal(t, gorm.ErrInvalidValue, err)
}

func TestTaskRepositoryIsExistsByIdExisted(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	taskId, err := uuid.Parse(faker.UUIDHyphenated())

	require.NoError(t, err)

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(taskId))

	result := taskRepository.IsExistsById(taskId)

	require.True(t, result)
}

func TestTaskRepositoryIsExistsByIdNotExisted(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	mock.ExpectQuery("SELECT").WillReturnError(gorm.ErrRecordNotFound)

	result := taskRepository.IsExistsById(taskId)

	require.False(t, result)
}

func TestTaskRepositoryGetAllByUserIdSuccess(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	userId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(taskId))

	result := taskRepository.GetAllByUserId(userId)

	require.IsType(t, &[]db.Task{}, result)
	require.NotEmpty(t, result)
	tasks := *result
	require.Equal(t, taskId, tasks[0].ID)
}

func TestTaskRepositoryGetAllByUserIdEmpty(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	userId, err := uuid.Parse(faker.UUIDHyphenated())

	require.NoError(t, err)

	taskRepository := repository.NewTaskRepository(mockedDatabase)

	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id"}))

	result := taskRepository.GetAllByUserId(userId)

	require.IsType(t, &[]db.Task{}, result)
	require.Empty(t, result)
}
