package service

import "bwastartup/model"

type ServiceUser interface {
	RegisterUser(input model.RegisterUserInput) (model.User, error)
	Login(input model.LoginInput) (model.User, error)
	IsEmailEvailable(input model.CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (model.User, error)
}
