package service_test

import (
	"errors"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"poymanov/todo/internal/service"
	mock_service "poymanov/todo/internal/service/mocks"
	"poymanov/todo/pkg/db"
	"poymanov/todo/pkg/jwt"
	"testing"
)

func TestAuthServiceRegister_UserAlreadyExists(t *testing.T) {
	authService, userService := mockAuthService(t)

	userService.EXPECT().FindByEmail(gomock.Any()).Return(&db.User{}, nil)

	token, err := authService.Register(service.RegisterData{})

	require.Empty(t, token)
	require.EqualError(t, err, service.ErrUserExists)
}

func TestAuthServiceRegister_Success(t *testing.T) {
	authService, userService := mockAuthService(t)

	userService.EXPECT().FindByEmail(gomock.Any()).Return(nil, errors.New(faker.Word()))
	userService.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(&db.User{}, nil)

	token, err := authService.Register(service.RegisterData{})

	require.NoError(t, err)
	require.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiJ9.hLLtvTt3MOzBg4X_PChuRTaGvNP_P68YJ83I-jKewhw", token)
}

func TestAuthServiceLogin_NotExistedUser(t *testing.T) {
	authService, userService := mockAuthService(t)

	userService.EXPECT().FindByEmail(gomock.Any()).Return(nil, errors.New(faker.Word()))

	token, err := authService.Login(service.LoginData{})

	require.Empty(t, token)
	require.EqualError(t, err, service.ErrWrongCredentials)
}

func TestAuthServiceLogin_WrongPassword(t *testing.T) {
	authService, userService := mockAuthService(t)

	userService.EXPECT().FindByEmail(gomock.Any()).Return(&db.User{}, nil)

	token, err := authService.Login(service.LoginData{})

	require.Empty(t, token)
	require.EqualError(t, err, service.ErrWrongCredentials)
}

func TestAuthServiceLogin_Success(t *testing.T) {
	authService, userService := mockAuthService(t)

	userService.EXPECT().FindByEmail(gomock.Any()).Return(&db.User{
		Password: "$2a$10$RxUZBWvGvCOXWQvI2QWpeuL6f3aksSdTQtOkG2TglZkqV4jbTGlwm",
	}, nil)

	token, err := authService.Login(service.LoginData{Password: "123qwe"})

	require.NoError(t, err)
	require.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiJ9.hLLtvTt3MOzBg4X_PChuRTaGvNP_P68YJ83I-jKewhw", token)
}

func mockAuthService(t *testing.T) (*service.AuthService, *mock_service.MockUser) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	userService := mock_service.NewMockUser(mockCtl)

	jwtHelper := jwt.NewJWT(faker.JWT)

	authService := service.NewAuthService(userService, jwtHelper)

	return authService, userService
}
