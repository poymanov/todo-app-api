package main

import (
	"fmt"
	"net/http"
	"poymanov/todo/internal/healthcheck"
)

func App() http.Handler {
	router := http.NewServeMux()

	healthcheck.NewHealthCheckHandler(router)

	return router
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	fmt.Println("Сервер запущен")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
