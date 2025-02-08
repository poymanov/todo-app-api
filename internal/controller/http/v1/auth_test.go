package v1

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"poymanov/todo/internal/service"
	mock_service "poymanov/todo/internal/service/mocks"
	"testing"
)

func TestAuthRegister(t *testing.T) {
	testCases := []struct {
		name         string
		body         string
		response     string
		statusCode   int
		mockFunction func(authService *mock_service.MockAuth)
	}{
		{
			name:         "Empty",
			body:         ``,
			response:     `{"message":"EOF"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(authService *mock_service.MockAuth) {},
		},
		{
			name:         "Missing name",
			body:         `{"email": "test@test.com", "password": "test"}`,
			response:     `{"message":"Key: 'RegisterRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(authService *mock_service.MockAuth) {},
		},
		{
			name:         "Missing email",
			body:         `{"name": "test", "password": "test"}`,
			response:     `{"message":"Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(authService *mock_service.MockAuth) {},
		},
		{
			name:         "Wrong email",
			body:         `{"name": "test", "email": "test" , "password": "test"}`,
			response:     `{"message":"Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(authService *mock_service.MockAuth) {},
		},
		{
			name:         "Missing password",
			body:         `{"name": "test", "email": "test@test.com"}`,
			response:     `{"message":"Key: 'RegisterRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(authService *mock_service.MockAuth) {},
		},
		{
			name:       "Success",
			body:       `{"name": "test","email": "test@test.com", "password": "test"}`,
			response:   `{"token":"token"}`,
			statusCode: http.StatusOK,
			mockFunction: func(authService *mock_service.MockAuth) {
				authService.EXPECT().Register(gomock.Any()).Return("token", nil).AnyTimes()
			},
		},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			authService := mock_service.NewMockAuth(c)
			tc.mockFunction(authService)
			handler := Handler{services: &service.Services{Auth: authService}}

			r := gin.New()
			r.POST("/auth/register", handler.register)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/register", bytes.NewBufferString(tc.body))
			r.ServeHTTP(w, req)

			require.Equal(t, w.Code, tc.statusCode)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}

func TestAuthLogin(t *testing.T) {
	testCases := []struct {
		name         string
		body         string
		response     string
		statusCode   int
		mockFunction func(authService *mock_service.MockAuth)
	}{
		{
			name:         "Empty",
			body:         ``,
			response:     `{"message":"EOF"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(authService *mock_service.MockAuth) {},
		},
		{
			name:         "Missing email",
			body:         `{"password": "test"}`,
			response:     `{"message":"Key: 'LoginRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(authService *mock_service.MockAuth) {},
		},
		{
			name:         "Wrong email",
			body:         `{"email": "test" , "password": "test"}`,
			response:     `{"message":"Key: 'LoginRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(authService *mock_service.MockAuth) {},
		},
		{
			name:         "Missing password",
			body:         `{"email": "test@test.com"}`,
			response:     `{"message":"Key: 'LoginRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
			statusCode:   http.StatusUnprocessableEntity,
			mockFunction: func(authService *mock_service.MockAuth) {},
		},
		{
			name:       "Success",
			body:       `{"email": "test@test.com", "password": "test"}`,
			response:   `{"token":"token"}`,
			statusCode: http.StatusOK,
			mockFunction: func(authService *mock_service.MockAuth) {
				authService.EXPECT().Login(gomock.Any()).Return("token", nil).AnyTimes()
			},
		},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			authService := mock_service.NewMockAuth(c)
			tc.mockFunction(authService)
			handler := Handler{services: &service.Services{Auth: authService}}

			r := gin.New()
			r.POST("/auth/login", handler.login)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/login", bytes.NewBufferString(tc.body))

			// Make Request
			r.ServeHTTP(w, req)

			require.Equal(t, w.Code, tc.statusCode)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}
