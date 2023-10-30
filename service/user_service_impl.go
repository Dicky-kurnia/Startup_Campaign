package service

import (
	"bwastartup/model"
	"bwastartup/repository"
	"errors"

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

func (service *serviceUser) Login(input model.LoginInput) (model.User, error) {
	email := input.Email
	password := input.Password

	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *serviceUser) IsEmailEvailable(input model.CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (service *serviceUser) SaveAvatar(ID int, fileLocation string) (model.User, error) {
	user, err := service.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := service.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}
