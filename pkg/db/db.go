package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"poymanov/todo/config"
	"time"
)

func NewDb(conf *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(conf.DB.DbConnectionAsString()), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		panic(err)
	}

	return db
}
