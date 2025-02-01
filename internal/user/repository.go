package user

import (
	"gorm.io/gorm"
	"poymanov/todo/pkg/db"
)

type UserRepository struct {
	Db *gorm.DB
}

type UserRepositoryDeps struct {
	Db *gorm.DB
}

func NewUserRepository(deps UserRepositoryDeps) *UserRepository {
	return &UserRepository{Db: deps.Db}
}

func (repo *UserRepository) Create(user *db.User) (*db.User, error) {
	result := repo.Db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*db.User, error) {
	var user db.User
	result := repo.Db.First(&user, "email=?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
