package profile

import (
	"errors"
	"net/http"
	"poymanov/todo/internal/service"
	"poymanov/todo/pkg/jwt"
	"poymanov/todo/pkg/middleware"
	"poymanov/todo/pkg/response"
)

const ErrFailedToGetProfile = "failed to get profile"

type Profile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileHandler struct {
	services *service.Services
	jwt      *jwt.JWT
}

func NewProfileHandler(router *http.ServeMux, services *service.Services, jwt *jwt.JWT) {
	handler := &ProfileHandler{
		services: services,
		jwt:      jwt,
	}
	router.Handle("GET /profile", middleware.Auth(handler.getProfile(), jwt))
}

// @Description	Получение профиля текущего авторизованного пользователя
// @Tags			profile
// @Success		200	{object}	profile.Profile
// @Failure		400	{object}	response.ErrorResponse
// @Security		ApiKeyAuth
// @Router			/profile [get]
func (h *ProfileHandler) getProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		email, ok := req.Context().Value(middleware.ContextEmailKey).(string)

		if !ok {
			response.JsonError(w, errors.New(ErrFailedToGetProfile), http.StatusBadRequest)
		}

		existedUser, _ := h.services.User.FindByEmail(email)

		if existedUser == nil {
			response.JsonError(w, errors.New(ErrFailedToGetProfile), http.StatusBadRequest)
			return
		}

		profileResponse := &Profile{
			ID:    existedUser.ID.String(),
			Name:  existedUser.Name,
			Email: existedUser.Email,
		}

		response.Json(w, profileResponse, http.StatusOK)
	}
}
