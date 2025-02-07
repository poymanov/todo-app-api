package main

import (
	"poymanov/todo/config"
	"poymanov/todo/internal/domain"
	"poymanov/todo/pkg/db"
)

func main() {
	conf := config.NewConfig()
	database := db.NewDb(conf)

	err := database.AutoMigrate(&domain.User{}, &domain.Task{})

	if err != nil {
		panic(err)
	}
}
