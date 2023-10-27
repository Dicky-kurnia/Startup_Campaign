package service

import (
	"bwastartup/model"
	"bwastartup/repository"

	"golang.org/x/crypto/bcrypt"
)

type serviceUser struct {
	repository repository.RepositoryUser
}

func NewServiceUser(repository repository.RepositoryUser) ServiceUser {
	return &serviceUser{repository: repository}
}

func (service *serviceUser) RegisterUser(input model.RegisterUserInput) (model.User, error) {
	user := model.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := service.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
