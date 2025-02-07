package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poymanov/todo/pkg/response"
)

const ErrFailedToGetProfile = "failed to get profile"

type Profile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *Handler) initProfileRoutes(api *gin.RouterGroup) {
	api.GET("/profile", h.auth, h.getProfile)
}

// @Description	Получение профиля текущего авторизованного пользователя
// @Tags			profile
// @Success		200	{object}	Profile
// @Failure		400	{object}	response.ErrorResponse
// @Security		ApiKeyAuth
// @Router			/profile [get]
func (h *Handler) getProfile(c *gin.Context) {
	email, err := getContextEmail(c)

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToGetProfile)
		return
	}

	existedUser, _ := h.services.User.FindByEmail(email)

	if existedUser == nil {
		response.NewErrorResponse(c, http.StatusBadRequest, ErrFailedToGetProfile)
		return
	}

	profileResponse := &Profile{
		ID:    existedUser.ID.String(),
		Name:  existedUser.Name,
		Email: existedUser.Email,
	}

	c.JSON(http.StatusOK, profileResponse)
}
