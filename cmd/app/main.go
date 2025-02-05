package main

import (
	"fmt"
	"net/http"
	"poymanov/todo/config"
	"poymanov/todo/internal/auth"
	"poymanov/todo/internal/healthcheck"
	"poymanov/todo/internal/profile"
	"poymanov/todo/internal/repository"
	"poymanov/todo/internal/service"
	"poymanov/todo/internal/swagger"
	"poymanov/todo/internal/task"
	"poymanov/todo/pkg/db"
	"poymanov/todo/pkg/jwt"
)

// @title						To-Do App API
// @version					1.0
// @description				API приложения для ведения списка дел
// @contact.name				Николай Пойманов
// @contact.email				n.poymanov@gmail.com
//
// @host						localhost:8099
// @BasePath					/
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				JWT-токен, в формате `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5ydSJ9.QiiLTDNqzID55nlQnYgmminveyKs2kzbwnGCEQqyc1A`
func App() http.Handler {
	conf := config.NewConfig()
	database := db.NewDb(conf)

	router := http.NewServeMux()

	// Common
	jwtHelper := jwt.NewJWT(conf.Auth.Secret)

	// Repositories
	repositories := repository.NewRepositories(database)
	services := service.NewServices(repositories, jwtHelper)

	// Handlers
	healthcheck.NewHealthCheckHandler(router)
	auth.NewAuthHandler(router, services)
	profile.NewProfileHandler(router, services, jwtHelper)
	task.NewTaskHandler(router, services, jwtHelper)
	swagger.NewSwaggerHandler(router)

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
