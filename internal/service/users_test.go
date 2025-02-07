package service_test

import (
	"errors"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mock_repository "poymanov/todo/internal/repository/mocks"
	"poymanov/todo/internal/service"
	"poymanov/todo/pkg/db"
	"testing"
)

func TestUserServiceCreate_Failed(t *testing.T) {
	userService, userRepo := mockUserService(t)

	userRepo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("failed"))

	user, err := userService.Create(faker.Name(), faker.Email(), faker.Password())

	require.Error(t, err)
	require.Nil(t, user)
}

func TestUserServiceCreate_Success(t *testing.T) {
	userService, userRepo := mockUserService(t)

	userData := db.User{Name: faker.Name(), Email: faker.Email()}

	userRepo.EXPECT().Create(gomock.Any()).Return(&userData, nil)

	createdUser, err := userService.Create(userData.Name, userData.Email, faker.Password())

	require.NoError(t, err)
	require.NotNil(t, createdUser)
	require.Equal(t, userData.Email, createdUser.Email)
	require.Equal(t, userData.Name, createdUser.Name)
}

func TestUserServiceFindByEmail_Failed(t *testing.T) {
	userService, userRepo := mockUserService(t)

	userRepo.EXPECT().FindByEmail(gomock.Any()).Return(nil, errors.New("failed"))

	userFind, err := userService.FindByEmail(faker.Email())

	require.Error(t, err)
	require.Nil(t, userFind)
}

func TestUserServiceFindByEmail_Success(t *testing.T) {
	userService, userRepo := mockUserService(t)

	userData := db.User{Name: faker.Name(), Email: faker.Email()}

	userRepo.EXPECT().FindByEmail(gomock.Any()).Return(&userData, nil)

	userFind, err := userService.FindByEmail(faker.Email())

	require.NoError(t, err)
	require.NotNil(t, userFind)
	require.Equal(t, userData.Email, userFind.Email)
	require.Equal(t, userData.Name, userFind.Name)
}

func mockUserService(t *testing.T) (*service.UserService, *mock_repository.MockUser) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	userRepo := mock_repository.NewMockUser(mockCtl)

	authService := service.NewUserService(userRepo)

	return authService, userRepo
}
