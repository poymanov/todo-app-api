package v1

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"poymanov/todo/internal/service"
	mock_service "poymanov/todo/internal/service/mocks"
	"testing"
)

func TestAuthRegister_RequestBodyValidaton(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		response string
	}{
		{
			name:     "Empty",
			body:     ``,
			response: `{"message":"EOF"}`,
		},
		{
			name:     "Missing name",
			body:     `{"email": "test@test.com", "password": "test"}`,
			response: `{"message":"Key: 'RegisterRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
		{
			name:     "Missing email",
			body:     `{"name": "test", "password": "test"}`,
			response: `{"message":"Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
		},
		{
			name:     "Wrong email",
			body:     `{"name": "test", "email": "test" , "password": "test"}`,
			response: `{"message":"Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`,
		},
		{
			name:     "Missing password",
			body:     `{"name": "test", "email": "test@test.com"}`,
			response: `{"message":"Key: 'RegisterRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
	}

	handler := Handler{services: &service.Services{}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.New()
			r.POST("/auth/register", handler.register)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/register", bytes.NewBufferString(tc.body))
			r.ServeHTTP(w, req)

			require.Equal(t, w.Code, http.StatusUnprocessableEntity)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}

func TestAuthRegister_Success(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	token := faker.Jwt()

	authService := mock_service.NewMockAuth(c)
	authService.EXPECT().Register(gomock.Any()).Return(token, nil).AnyTimes()
	handler := Handler{services: &service.Services{Auth: authService}}

	r := gin.New()
	r.POST("/auth/register", handler.register)

	requestBody := `{"name": "test","email": "test@test.com", "password": "test"}`

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBufferString(requestBody))
	r.ServeHTTP(w, req)

	response := fmt.Sprintf(`{"token":"%s"}`, token)

	require.Equal(t, w.Code, http.StatusOK)
	require.Equal(t, response, w.Body.String())
}

func TestAuthLogin_RequestBodyValidaton(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		response string
	}{
		{
			name:     "Empty",
			body:     ``,
			response: `{"message":"EOF"}`,
		},
		{
			name:     "Missing email",
			body:     `{"name": "test", "password": "test"}`,
			response: `{"message":"Key: 'LoginRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
		},
		{
			name:     "Wrong email",
			body:     `{"name": "test", "email": "test" , "password": "test"}`,
			response: `{"message":"Key: 'LoginRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`,
		},
		{
			name:     "Missing password",
			body:     `{"name": "test", "email": "test@test.com"}`,
			response: `{"message":"Key: 'LoginRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
	}

	handler := Handler{services: &service.Services{}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.New()
			r.POST("/auth/login", handler.login)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/login", bytes.NewBufferString(tc.body))

			// Make Request
			r.ServeHTTP(w, req)

			require.Equal(t, w.Code, http.StatusUnprocessableEntity)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}

func TestAuthLogin_Success(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	token := faker.Jwt()

	authService := mock_service.NewMockAuth(c)
	authService.EXPECT().Login(gomock.Any()).Return(token, nil).AnyTimes()
	handler := Handler{services: &service.Services{Auth: authService}}

	r := gin.New()
	r.POST("/auth/login", handler.login)

	requestBody := `{"email": "test@test.com", "password": "test"}`

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewBufferString(requestBody))
	r.ServeHTTP(w, req)

	response := fmt.Sprintf(`{"token":"%s"}`, token)

	require.Equal(t, w.Code, http.StatusOK)
	require.Equal(t, response, w.Body.String())
}
