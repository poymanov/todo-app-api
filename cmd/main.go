package main

import (
	"fmt"
	"net/http"
	"poymanov/todo/config"
	"poymanov/todo/internal/auth"
	"poymanov/todo/internal/healthcheck"
	"poymanov/todo/internal/user"
	"poymanov/todo/pkg/db"
	jwt2 "poymanov/todo/pkg/jwt"
)

func App() http.Handler {
	conf := config.NewConfig()
	database := db.NewDb(conf)

	router := http.NewServeMux()

	// Common
	jwt := jwt2.NewJWT(conf.Auth.Secret)

	// Repositories
	userRepository := user.UserRepository{Db: database}

	// Services
	userService := user.NewUserService(user.UserServiceDeps{UserRepository: userRepository})
	authService := auth.NewAuthService(auth.AuthServiceDeps{UserService: userService, JWT: jwt})

	// Handlers
	healthcheck.NewHealthCheckHandler(router)
	auth.NewAuthHandlerHandler(router, auth.AuthHandlerDeps{AuthService: authService})

	return router
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8080")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
