package healthcheck

import (
	"net/http"
	"poymanov/todo/pkg/response"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler(router *http.ServeMux) {
	handler := &HealthCheckHandler{}
	router.HandleFunc("GET /healthcheck", handler.CurrentStatus())
}

func (h *HealthCheckHandler) CurrentStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		healthCheckResponse := &HealthCheckResponse{
			Status: "ok",
		}

		response.Json(w, healthCheckResponse, http.StatusOK)
	}
}
