package app

import (
	"fmt"
	"net/http"
	"poymanov/todo/config"
	controllerHandler "poymanov/todo/internal/controller/http"
	"poymanov/todo/internal/repository"
	"poymanov/todo/internal/service"
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
// @BasePath					/api/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				JWT-токен, в формате `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5ydSJ9.QiiLTDNqzID55nlQnYgmminveyKs2kzbwnGCEQqyc1A`
func Run() {
	conf := config.NewConfig()
	database := db.NewDb(conf)

	jwtHelper := jwt.NewJWT(conf.Auth.Secret)

	repositories := repository.NewRepositories(database)
	services := service.NewServices(repositories, jwtHelper)

	handler := controllerHandler.NewHandler(services, jwtHelper)
	router := handler.Init()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8080")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
