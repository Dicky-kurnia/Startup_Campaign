package repository

import (
	"bwastartup/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) RepositoryUser {
	return &userRepository{db: db}
}

func (repository *userRepository) Save(user model.User) (model.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User

	err := repository.db.Where("email: ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
