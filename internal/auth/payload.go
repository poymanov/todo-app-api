package auth

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type RegisterData struct {
	Name     string
	Email    string
	Password string
}
