package main

import (
	"poymanov/todo/config"
	"poymanov/todo/internal/task"
	"poymanov/todo/internal/user"
	"poymanov/todo/pkg/db"
)

func main() {
	conf := config.NewConfig()
	database := db.NewDb(conf)

	err := database.AutoMigrate(&user.User{}, &task.Task{})

	if err != nil {
		panic(err)
	}
}
