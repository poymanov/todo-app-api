package profile

import (
	"errors"
	"net/http"
	"poymanov/todo/internal/user"
	"poymanov/todo/pkg/jwt"
	"poymanov/todo/pkg/middleware"
	"poymanov/todo/pkg/response"
)

type ProfileHandlerDeps struct {
	UserService *user.UserService
	JWT         *jwt.JWT
}

type ProfileHandler struct {
	UserService *user.UserService
}

func NewProfileHandler(router *http.ServeMux, deps ProfileHandlerDeps) {
	handler := &ProfileHandler{
		UserService: deps.UserService,
	}
	router.Handle("GET /profile", middleware.Auth(handler.getProfile(), deps.JWT))
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

		existedUser, _ := h.UserService.FindByEmail(email)

		if existedUser == nil {
			response.JsonError(w, errors.New(ErrFailedToGetProfile), http.StatusBadRequest)
		}

		profileResponse := &Profile{
			ID:    existedUser.ID.String(),
			Name:  existedUser.Name,
			Email: existedUser.Email,
		}

		response.Json(w, profileResponse, http.StatusOK)
	}
}
