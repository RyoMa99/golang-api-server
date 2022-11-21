package persistence

import (
	"github.com/RyoMa99/go_ddd/domain/model"
	"github.com/RyoMa99/go_ddd/domain/repository"
	"github.com/jinzhu/gorm"
)

// UserにおけるPersistenceのインターフェース
type userPersistence struct {
	Conn *gorm.DB
}

// Userデータに関するPersistenceを生成
func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

// 検索
func (up *userPersistence) Search(name string) ([]*model.User, error) {
	var user []*model.User

	if name == "" {
		// nameが空なら全件検索
		if err := up.Conn.Find(&user).Error; err != nil {
			return nil, err
		}
	} else {
		// nameでの条件付き検索
		if err := up.Conn.Where("name = ?", name).Find(&user).Error; err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (up *userPersistence) Create(user *model.User) (*model.User, error) {
	if result := up.Conn.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
