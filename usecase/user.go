package usecase

import (
	"github.com/RyoMa99/go_ddd/domain/model"
	"github.com/RyoMa99/go_ddd/domain/repository"
)

// UserにおけるUseCaseのインターフェース
type UserUseCase interface {
	Search(name string) ([]*model.User, error)
	Create(name string) (*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// Userデータに関するUseCaseを生成
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

// 検索
func (uu userUseCase) Search(name string) (user []*model.User, err error) {
	user, err = uu.userRepository.Search(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 新規作成
func (uu userUseCase) Create(name string) (*model.User, error) {
	user, err := model.NewUser(name)
	if err != nil {
		return nil, err
	}

	createdUser, err := uu.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
