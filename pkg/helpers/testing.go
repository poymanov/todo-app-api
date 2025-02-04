package helpers

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMockDatabase() (*gorm.DB, sqlmock.Sqlmock) {
	database, mock, err := sqlmock.New()

	if err != nil {
		panic(err)
	}

	mockedDatabase, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	return mockedDatabase, mock
}
