package auth

import (
	"net/http"
	"poymanov/todo/pkg/request"
	"poymanov/todo/pkg/response"
)

type AuthHandler struct {
}

func NewAuthHandlerHandler(router *http.ServeMux) {
	handler := &AuthHandler{}
	router.HandleFunc("POST /auth/register", handler.register())
}

func (h *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](req)

		if err != nil {
			response.JsonError(w, err, http.StatusUnprocessableEntity)
			return
		}

		response.Json(w, body, http.StatusOK)
	}
}
