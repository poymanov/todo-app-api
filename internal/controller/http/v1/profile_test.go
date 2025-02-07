package v1

import (
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
)

func TestGetProfile_FailedToGetEmailFromContext(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	handler := Handler{services: &service.Services{}}

	r := gin.New()
	r.GET("/profile", handler.getProfile)

	// Create Request
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/profile", nil)

	// Make Request
	r.ServeHTTP(w, req)

	require.Equal(t, w.Code, http.StatusBadRequest)
	require.Equal(t, `{"message":"Failed to get profile"}`, w.Body.String())
}

func TestGetProfile_NotExistedUser(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	userService := mock_service.NewMockUser(c)
	userService.EXPECT().FindByEmail(gomock.Any()).Return(nil, errors.New("failed"))

	handler := Handler{services: &service.Services{User: userService}}

	r := gin.New()
	r.GET("/profile", func(c *gin.Context) {
		c.Set(ContextEmailKey, faker.Email())
	}, handler.getProfile)

	// Create Request
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/profile", nil)

	// Make Request
	r.ServeHTTP(w, req)

	require.Equal(t, w.Code, http.StatusBadRequest)
	require.Equal(t, `{"message":"Failed to get profile"}`, w.Body.String())
}

func TestGetProfile_Success(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	userId, err := uuid.Parse(faker.UUIDHyphenated())
	require.NoError(t, err)

	user := domain.User{ID: userId, Email: faker.Email(), Name: faker.Name()}

	userService := mock_service.NewMockUser(c)
	userService.EXPECT().FindByEmail(gomock.Any()).Return(&user, nil)

	handler := Handler{services: &service.Services{User: userService}}

	r := gin.New()
	r.GET("/profile", func(c *gin.Context) {
		c.Set(ContextEmailKey, faker.Email())
	}, handler.getProfile)

	// Create Request
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/profile", nil)

	// Make Request
	r.ServeHTTP(w, req)

	expectedMessage := fmt.Sprintf(`{"id":"%s","name":"%s","email":"%s"}`, user.ID.String(), user.Name, user.Email)

	require.Equal(t, w.Code, http.StatusOK)
	require.Equal(t, expectedMessage, w.Body.String())
}
