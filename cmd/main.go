package main

import (
	"fmt"
	"net/http"
	"poymanov/todo/config"
	"poymanov/todo/internal/auth"
	"poymanov/todo/internal/healthcheck"
	"poymanov/todo/internal/profile"
	"poymanov/todo/internal/swagger"
	"poymanov/todo/internal/task"
	"poymanov/todo/internal/user"
	"poymanov/todo/pkg/db"
	"poymanov/todo/pkg/jwt"
)

//	@title			TO-DO App API
//	@version		1.0
//	@description	API приложения для ведения списка дел
//	@contact.name	Николай Пойманов
//	@contact.email	n.poymanov@gmail.com
//
// @host		localhost:8099
// @BasePath	/
func App() http.Handler {
	conf := config.NewConfig()
	database := db.NewDb(conf)

	router := http.NewServeMux()

	// Common
	jwtHelper := jwt.NewJWT(conf.Auth.Secret)

	// Repositories
	userRepository := user.NewUserRepository(user.UserRepositoryDeps{Db: database})
	taskRepository := task.NewTaskRepository(task.TaskRepositoryDeps{Db: database})

	// Services
	userService := user.NewUserService(user.UserServiceDeps{UserRepository: userRepository})
	authService := auth.NewAuthService(auth.AuthServiceDeps{UserService: userService, JWT: jwtHelper})
	taskService := task.NewTaskService(task.TaskServiceDeps{TaskRepository: taskRepository})

	// Handlers
	healthcheck.NewHealthCheckHandler(router)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{AuthService: authService})
	profile.NewProfileHandler(router, profile.ProfileHandlerDeps{JWT: jwtHelper, UserService: userService})
	task.NewTaskHandler(router, task.TaskHandlerDeps{JWT: jwtHelper, UserService: userService, TaskService: taskService})
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
