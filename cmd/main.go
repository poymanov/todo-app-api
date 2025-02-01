package main

import (
	"fmt"
	"net/http"
	"poymanov/todo/internal/auth"
	"poymanov/todo/internal/healthcheck"
)

func App() http.Handler {
	//conf := config.NewConfig()
	//database := db.NewDb(conf)

	router := http.NewServeMux()

	healthcheck.NewHealthCheckHandler(router)
	auth.NewAuthHandlerHandler(router)

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
