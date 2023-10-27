package service

import "bwastartup/model"

type ServiceUser interface {
	RegisterUser(input model.RegisterUserInput) (model.User, error)
}
