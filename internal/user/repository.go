package user

import "gorm.io/gorm"

type UserRepository struct {
	Db *gorm.DB
}

type UserRepositoryDeps struct {
	Db *gorm.DB
}

func NewUserRepository(deps UserRepositoryDeps) *UserRepository {
	return &UserRepository{Db: deps.Db}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.Db.First(&user, "email=?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
