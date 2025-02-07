package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poymanov/todo/internal/service"
	"poymanov/todo/pkg/response"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	api.POST("/auth/register", h.register)
	api.POST("/auth/login", h.login)
}

// @Description	Регистрация пользователя
// @Tags			auth
// @Param			register	body		RegisterRequest	true	"Данные нового пользователя"
// @Success		201			{object}	RegisterResponse
// @Failure		400			{object}	response.ErrorResponse
// @Failure		422			{object}	response.ErrorResponse
// @Router			/auth/register [post]
func (h *Handler) register(c *gin.Context) {
	var body RegisterRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	token, err := h.services.Auth.Register(service.RegisterData{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{Token: token})
}

// @Description	Авторизация пользователя
// @Tags			auth
// @Param			login	body		LoginRequest	true	"Данные зарегистрированного  пользователя"
// @Success		200		{object}	LoginResponse
// @Failure		400		{object}	response.ErrorResponse
// @Failure		422		{object}	response.ErrorResponse
// @Router			/auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var body LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		response.NewErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	token, err := h.services.Auth.Login(service.LoginData{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
