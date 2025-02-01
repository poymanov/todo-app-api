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

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.register())
	router.HandleFunc("POST /auth/login", handler.login())
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

		response.Json(w, RegisterResponse{Token: token}, http.StatusCreated)
	}
}

func (h *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LoginRequest](req)

		if err != nil {
			response.JsonError(w, err, http.StatusUnprocessableEntity)
			return
		}

		token, err := h.AuthService.Login(LoginData{
			Email:    body.Email,
			Password: body.Password,
		})

		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := LoginResponse{
			Token: token,
		}

		response.Json(w, res, http.StatusOK)
	}
}
