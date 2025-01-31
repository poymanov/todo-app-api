package main

import (
	"poymanov/todo/config"
	"poymanov/todo/internal/user"
	"poymanov/todo/pkg/db"
)

func main() {
	conf := config.NewConfig()
	database := db.NewDb(conf)

	err := database.AutoMigrate(&user.User{})

	if err != nil {
		panic(err)
	}
}
