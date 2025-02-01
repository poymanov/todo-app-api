package auth

import (
	"net/http"
	"poymanov/todo/pkg/request"
	"poymanov/todo/pkg/response"
)

type AuthHandlerDeps struct {
	AuthService *AuthService
}

type AuthHandler struct {
	AuthService *AuthService
}

func NewAuthHandlerHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.register())
}

func (h *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](req)

		if err != nil {
			response.JsonError(w, err, http.StatusUnprocessableEntity)
			return
		}

		token, err := h.AuthService.Register(RegisterData{
			Name:     body.Name,
			Email:    body.Email,
			Password: body.Password,
		})

		if err != nil {
			response.JsonError(w, err, http.StatusBadRequest)
			return
		}

		response.Json(w, RegisterResponse{Token: token}, http.StatusOK)
	}
}
