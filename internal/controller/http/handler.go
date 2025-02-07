package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "poymanov/todo/docs"
	v1 "poymanov/todo/internal/controller/http/v1"
	"poymanov/todo/internal/service"
	"poymanov/todo/pkg/jwt"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

type Handler struct {
	services *service.Services
	jwt      *jwt.JWT
}

func NewHandler(services *service.Services, jwt *jwt.JWT) *Handler {
	return &Handler{services: services, jwt: jwt}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	initSwaggerRoute(router)
	initHealthCheck(router)

	h.initAPI(router)

	return router
}

func initSwaggerRoute(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @Description	Получение статуса работоспособности приложения
// @Tags			common
// @Success		200	{object}	HealthCheckResponse
//
// @Router			/healthcheck [get]
func initHealthCheck(router *gin.Engine) {
	healthCheckResponse := &HealthCheckResponse{
		Status: "ok",
	}

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthCheckResponse)
	})
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.jwt)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
