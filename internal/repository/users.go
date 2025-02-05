package repository

import (
	"gorm.io/gorm"
	"poymanov/todo/pkg/db"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(user *db.User) (*db.User, error) {
	result := repo.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*db.User, error) {
	var user db.User
	result := repo.db.First(&user, "email=?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
