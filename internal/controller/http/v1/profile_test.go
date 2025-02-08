package v1

import (
	"errors"
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
)

func TestGetProfile(t *testing.T) {
	testCases := []struct {
		name            string
		response        string
		statusCode      int
		contextModifier func(c *gin.Context)
		mockFunction    func(userService *mock_service.MockUser)
	}{
		{
			name:            "Failed to get email from context",
			response:        `{"message":"Failed to get profile"}`,
			statusCode:      http.StatusBadRequest,
			mockFunction:    func(userService *mock_service.MockUser) {},
			contextModifier: func(c *gin.Context) {},
		},
		{
			name:       "Not existed user",
			response:   `{"message":"Failed to get profile"}`,
			statusCode: http.StatusBadRequest,
			mockFunction: func(userService *mock_service.MockUser) {
				userService.EXPECT().FindByEmail(gomock.Any()).Return(nil, errors.New("failed"))
			},
			contextModifier: func(c *gin.Context) {
				c.Set(ContextEmailKey, faker.Email())
			},
		},
		{
			name:       "Success",
			response:   `{"id":"64f7ecf1-cf5d-4f7f-888b-f3b68b68e70b","name":"test","email":"test@test.ru"}`,
			statusCode: http.StatusOK,
			mockFunction: func(userService *mock_service.MockUser) {
				userId, _ := uuid.Parse("64f7ecf1-cf5d-4f7f-888b-f3b68b68e70b")
				user := domain.User{ID: userId, Email: "test@test.ru", Name: "test"}
				userService.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)
			},
			contextModifier: func(c *gin.Context) {
				c.Set(ContextEmailKey, faker.Email())
			},
		},
	}

	c := gomock.NewController(t)
	defer c.Finish()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userService := mock_service.NewMockUser(c)
			tc.mockFunction(userService)
			handler := Handler{services: &service.Services{User: userService}}

			r := gin.New()
			r.GET("/profile", tc.contextModifier, handler.getProfile)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/profile", nil)

			r.ServeHTTP(w, req)

			require.Equal(t, w.Code, tc.statusCode)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}
