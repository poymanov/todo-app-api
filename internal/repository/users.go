package repository

import (
	"gorm.io/gorm"
	"poymanov/todo/internal/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(user *domain.User) (*domain.User, error) {
	result := repo.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := repo.db.First(&user, "email=?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
