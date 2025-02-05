package healthcheck

import (
	"net/http"
	"poymanov/todo/pkg/response"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

type HealthCheckHandler struct {
}

func NewHealthCheckHandler(router *http.ServeMux) {
	handler := &HealthCheckHandler{}
	router.HandleFunc("GET /healthcheck", handler.CurrentStatus())
}

// @Description	Получение статуса работоспособности приложения
// @Tags			common
// @Success		200	{object}	healthcheck.HealthCheckResponse
//
// @Router			/healthcheck [get]
func (h *HealthCheckHandler) CurrentStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		healthCheckResponse := &HealthCheckResponse{
			Status: "ok",
		}

		response.Json(w, healthCheckResponse, http.StatusOK)
	}
}
