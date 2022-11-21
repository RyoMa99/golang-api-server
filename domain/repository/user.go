package repository

import "github.com/RyoMa99/go_ddd/domain/model"

type UserRepository interface {
	Search(name string) ([]*model.User, error)
	Create(user *model.User) (*model.User, error)
}
