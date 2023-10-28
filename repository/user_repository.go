package repository

import "bwastartup/model"

type RepositoryUser interface {
	Save(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
}
