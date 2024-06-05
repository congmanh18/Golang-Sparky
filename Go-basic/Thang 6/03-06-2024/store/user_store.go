package store

import (
	"microlap-gorm/entity"

	"gorm.io/gorm"
)

type userStore struct {
	gorm *gorm.DB
}

func NewUserStore(gorm *gorm.DB) userStore {
	return userStore{
		gorm: gorm,
	}
}

func (u *userStore) Save(user entity.User) error {
	return u.gorm.Create(&user).Error
}

func (u *userStore) UpdateV2(id string, user entity.User) error {
	return u.gorm.Debug().Where("id = ?", id).Updates(&user).Error
}
