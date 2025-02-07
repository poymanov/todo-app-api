package v1

import (
	"github.com/gin-gonic/gin"
	"poymanov/todo/internal/service"
	"poymanov/todo/pkg/jwt"
)

type Handler struct {
	services *service.Services
	jwt      *jwt.JWT
}

func NewHandler(services *service.Services, jwt *jwt.JWT) *Handler {
	return &Handler{services: services, jwt: jwt}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initProfileRoutes(v1)
		h.initAuthRoutes(v1)
		h.initTasksRoutes(v1)
	}
}
