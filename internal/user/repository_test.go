package user_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"poymanov/todo/internal/user"
	"poymanov/todo/pkg/db"
	"poymanov/todo/pkg/helpers"
	"testing"
)

func TestUserRepositoryCreateSuccess(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	uuid := faker.UUIDHyphenated()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid))
	mock.ExpectCommit()

	userRepository := user.NewUserRepository(user.UserRepositoryDeps{Db: mockedDatabase})

	expectedUser := db.User{Email: faker.Email(), Name: faker.Name()}

	createdUser, err := userRepository.Create(&expectedUser)

	require.NoError(t, err)

	require.Equal(t, expectedUser.ID.String(), uuid)
	require.Equal(t, expectedUser.Name, createdUser.Name)
	require.Equal(t, expectedUser.Email, createdUser.Email)
}

func TestUserRepositoryCreateFailedEmailAlreadyExists(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnError(gorm.ErrDuplicatedKey)
	mock.ExpectRollback()

	userRepository := user.NewUserRepository(user.UserRepositoryDeps{Db: mockedDatabase})

	newUser := db.User{Email: faker.Email(), Name: faker.Name()}

	createdUser, err := userRepository.Create(&newUser)

	require.Nil(t, createdUser)
	require.Error(t, err)
	require.Equal(t, gorm.ErrDuplicatedKey, err)
}

func TestUserRepositoryFindByEmailSuccess(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	email := faker.Email()
	name := faker.Name()

	mockedDatabase.Create(&db.User{Email: email, Name: name})

	mock.ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow(faker.UUIDHyphenated(), name, email))

	userRepository := user.NewUserRepository(user.UserRepositoryDeps{Db: mockedDatabase})

	existedUser, err := userRepository.FindByEmail(email)

	require.NoError(t, err)

	require.Equal(t, name, existedUser.Name)
	require.Equal(t, email, existedUser.Email)
}

func TestUserRepositoryFindByEmailFailedNoExists(t *testing.T) {
	mockedDatabase, mock := helpers.InitMockDatabase()

	mock.ExpectQuery("SELECT").WillReturnError(gorm.ErrRecordNotFound)

	userRepository := user.NewUserRepository(user.UserRepositoryDeps{Db: mockedDatabase})

	existedUser, err := userRepository.FindByEmail(faker.Email())

	require.Nil(t, existedUser)
	require.Error(t, err)
	require.Equal(t, gorm.ErrRecordNotFound, err)
}
