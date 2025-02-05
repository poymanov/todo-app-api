package main

import (
	"poymanov/todo/config"
	"poymanov/todo/pkg/db"
)

func main() {
	conf := config.NewConfig()
	database := db.NewDb(conf)

	err := database.AutoMigrate(&db.User{}, &db.Task{})

	if err != nil {
		panic(err)
	}
}
